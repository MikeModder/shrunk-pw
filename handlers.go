package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teris-io/shortid"
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
	query := r.URL.Query()
	long := query.Get("url")

	newShort, _ := shortid.Generate()
	saveURL(newShort, long)
}

// "/unshorten?url=<short>"
//	short can be <id>
// Response:
//	{ short: "<short url>", full: "<original url>" }
func unshortenHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	short := query.Get("url")

	if hasFullURL(short) {
		json.NewEncoder(w).Encode(unshortenedResp{
			Success: true,
			Short:   short,
			Full:    getFullURL(short),
		})
	} else {
		json.NewEncoder(w).Encode(unshortenedResp{
			Success: false,
			Short:   short,
		})
	}
	return
}
