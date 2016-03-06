package gong

import (
	"net/http"
)

// Shim is an adapter from net/http to gong/scope
func Shim(fn func(s Scope)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fn(Scope{
			req: req,
			res: res,
		})
	}
}
