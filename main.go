package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	host = "0.0.0.0"
	port = 8080
)

func main() {
	log.Println("shrunk-pw starting...")
	log.Printf("Listening on %s:%d\n", host, port)

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/shorten", shortenHandler)
	r.HandleFunc("/unshorten", unshortenHandler)
	r.HandleFunc("/{id}", gotoUrlHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r))
}
