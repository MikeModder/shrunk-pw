package main

import "net/http"

// "/"
// Show homepage
func indexHandler(w http.ResponseWriter, r *http.Request) {

}

// "/{shortid}"
// Redirects to full URL
func gotoUrlHandler(w http.ResponseWriter, r *http.Request) {

}

// "/shorten?url=<url>"
// Response:
//	{ success: bool, msg: "failed <reason>/success", short: "<short url>" }
func shortenHandler(w http.ResponseWriter, r *http.Request) {

}

// "/unshorten?url=<short>"
//	short can be <id>
// Response:
//	{ short: "<short url>", full: "<original url>" }
func unshortenHandler(w http.ResponseWriter, r *http.Request) {

}
