package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// "/"
// Show homepage
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

// "/{shortid}"
// Redirects to full URL
func gotoURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["id"]

	if hasFullURL(short) {
		http.Redirect(w, r, getFullURL(short), http.StatusTemporaryRedirect)
	} else {
		fmt.Fprint(w, "that isn't a valid short code!")
	}
}

// "/shorten?url=<url>"
// Response:
//	{ success: bool, msg: "failed <reason>/success", short: "<short url>" }
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "shorten")
}

// "/unshorten?url=<short>"
//	short can be <id>
// Response:
//	{ short: "<short url>", full: "<original url>" }
func unshortenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unshorten")
}
