package gong

import (
	"net/http"
)

type Scope struct {
	req *http.Request
	res http.ResponseWriter
}

func (s *Scope) Write(out []byte) {
	s.res.Write(out)
}
