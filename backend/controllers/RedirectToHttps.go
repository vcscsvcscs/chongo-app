package controllers

import (
	"log"
	"net/http"
	"strings"
)

func RedirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request.
	host := strings.Split(r.Host, ":")[0]
	log.Println(r.Host)
	http.Redirect(w, r, "https://"+host+r.RequestURI, http.StatusMovedPermanently)
}
