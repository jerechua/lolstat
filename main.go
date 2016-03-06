package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"
	"./gong"
	"./handlers/home"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", gong.Shim(home.MainHandler))

	port := 8080
	log.Printf("Listening on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
