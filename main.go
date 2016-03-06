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
	r.HandleFunc("/", MainHandler)

	port := 8080
	log.Printf("Listening on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func MainHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("gorilla!\n"))
}
