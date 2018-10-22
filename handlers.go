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
	http.ServeFile(w, r, "templates/index.html")
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

	if long == "" {
		json.NewEncoder(w).Encode(shortenResp{
			Success: false,
			Message: "Failed to create short URL! No URL given!",
		})
		return
	}

	newShort, _ := shortid.Generate()
	if saveURL(newShort, long) {
		json.NewEncoder(w).Encode(shortenResp{
			Success: true,
			Message: "Success",
			Short:   fmt.Sprintf("%s/%s", appConfig.Domain, newShort),
		})
	} else {
		json.NewEncoder(w).Encode(shortenResp{
			Success: false,
			Message: "Failed to create short URL!",
		})
	}
}

// "/unshorten?url=<short>"
//	short can be <id>
// Response:
//	{ short: "<short url>", full: "<original url>" }
func unshortenHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	short := query.Get("url")

	if short == "" {
		json.NewEncoder(w).Encode(unshortenedResp{
			Success: false,
			Message: "Failed to find full URL! No URL given!",
		})
		return
	}

	if hasFullURL(short) {
		json.NewEncoder(w).Encode(unshortenedResp{
			Success: true,
			Short:   short,
			Message: "Success",
			Full:    getFullURL(short),
		})
	} else {
		json.NewEncoder(w).Encode(unshortenedResp{
			Success: false,
			Message: "Failed to get full URL! Is that a valid short code?",
			Short:   short,
		})
	}
	return
}
