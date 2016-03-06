package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"
	"./gong"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", gong.Shim(MainHandler))

	port := 8080
	log.Printf("Listening on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func MainHandler(scope gong.Scope) {
	scope.Write([]byte("gorilla!\n"))
}
