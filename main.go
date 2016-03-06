package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./db"
	"./gong"
	"./handlers/home"

	"github.com/gorilla/mux"
)

var (
	portFlag = flag.Int("port", 8080, "The port to start the server on.")
)

func init() {
	flag.Parse()
}

func main() {
	db.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", gong.Shim(home.MainHandler))

	log.Printf("Listening on port: %d", *portFlag)
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), r)
}
