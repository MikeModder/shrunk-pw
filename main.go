package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	appConfig config
)

func init() {
	log.Println("Loading and parsing config...")
	cfgFile, e := os.Open("config.json")
	if e != nil {
		log.Fatalf("[error] failed to open config.json: %s\n", e.Error())
	}
	cfgParser := json.NewDecoder(cfgFile)
	e = cfgParser.Decode(&appConfig)
	if e != nil {
		log.Fatalf("[error] failed to open config.json: %s\n", e.Error())
	}
}

func main() {
	log.Println("shrunk-pw starting...")
	log.Println("attempting to connect to redis...")

	setupDB()

	log.Printf("Listening on %s:%d\n", appConfig.Domain, appConfig.Port)

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/shorten", shortenHandler)
	r.HandleFunc("/unshorten", unshortenHandler)
	r.HandleFunc("/{id}", gotoURLHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", appConfig.Port), r))
}
