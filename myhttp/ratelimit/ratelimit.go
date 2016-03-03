package ratelimit

import (
	"sync"
	"time"

	"../../myhttp"
)

type RateLimitClient struct {
	mu     sync.RWMutex
	Limits []*Limit
}

func (rm *RateLimitClient) waitRequestAvailable() {
	available := make(chan bool)
	for _, l := range rm.Limits {
		go l.waitAvailable(available)
	}

	for i := 0; i < len(rm.Limits); i++ {
		// Wait for all limits too signal that we can send a request.
		<-available
	}
}

func (rm *RateLimitClient) applyRate() {
	for _, l := range rm.Limits {
		l.useRate()
	}
}

func (rm *RateLimitClient) Get(req *myhttp.Request, res chan *myhttp.Response, err chan error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.waitRequestAvailable()
	go req.GetAsync(res, err, func() { rm.applyRate() })
}

// Limit assumes there's only 1 request goroutine accessing it
type Limit struct {
	// Requests is the actual limit
	Requests int

	// Timeout is the interval before we can send another request.
	Timeout time.Duration

	mu   sync.Mutex
	used int
}

// recoverRate is a thread-safe recovery of the rate.
func (l *Limit) recoverRate() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.used--
}

// regenerate waits a certain amount of time then readds the rate limit.
func (l *Limit) regenerate() {
	time.Sleep(l.Timeout)
	l.recoverRate()
}

// useRate uses a rate limit and sets a timer to readd it.
func (l *Limit) useRate() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.used++
	go l.regenerate()
}

// hasAvailable checks to see if there are available rates
func (l *Limit) hasAvailable() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.used < l.Requests {
		return true
	}
	return false
}

// waitAvailable waits until a rate limit is available.
func (l *Limit) waitAvailable(available chan bool) {
	for {
		if l.hasAvailable() {
			available <- true
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
