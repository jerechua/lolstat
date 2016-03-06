package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", Shim(MainHandler))

	port := 8080
	log.Printf("Listening on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

type Scope struct {
	req *http.Request
	res http.ResponseWriter
}

func (s *Scope) Write(out []byte) {
	s.res.Write(out)
}

func Shim(fn func(scope Scope)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		s := Scope{
			req: req,
			res: res,
		}
		fn(s)
	}
}

func MainHandler(scope Scope) {
	scope.Write([]byte("gorilla!\n"))
}
