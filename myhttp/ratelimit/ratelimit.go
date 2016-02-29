package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

// limit assumes there's only 1 request goroutine accessing it
type limit struct {
	// requests is the actual limit
	requests int

	// t is the interval before we can send another request.
	t time.Duration

	mu   sync.Mutex
	used int
}

// recoverRate is a thread-safe recovery of the rate.
func (l *limit) recoverRate() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.used--
}

// regenerate waits a certain amount of time then readds the rate limit.
func (l *limit) regenerate() {
	time.Sleep(l.t)
	l.recoverRate()
}

// useRate uses a rate limit and sets a timer to readd it.
func (l *limit) useRate() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.used++
	go l.regenerate()
}

// hasAvailable checks to see if there are available rates
func (l *limit) hasAvailable() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.used < l.requests {
		return true
	}
	return false
}

// waitAvailable waits until a rate limit is available.
func (l *limit) waitAvailable(available chan bool) {
	for {
		if l.hasAvailable() {
			available <- true
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

type RequestMaker struct {
	mu     sync.RWMutex
	limits []*limit
}

func (rm *RequestMaker) waitRequestAvailable() {
	available := make(chan bool)
	for _, l := range rm.limits {
		go l.waitAvailable(available)
	}

	for i := 0; i < len(rm.limits); i++ {
		// Wait for all limits too signal that we can send a request.
		<-available
	}
}

func (rm *RequestMaker) applyRate() {
	for _, l := range rm.limits {
		l.useRate()
	}
}

func (rm *RequestMaker) MakeRequest(s string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.waitRequestAvailable()
	fmt.Println(s)
	rm.applyRate()
}
