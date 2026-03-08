package main

import (
	"net/http"
	"os"
)

func main() {
	target := os.Getenv("REDIRECT_TARGET")
	if target == "" {
		target = "https://joinheritable.com"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		newURL := target + r.URL.Path
		if r.URL.RawQuery != "" {
			newURL += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, newURL, http.StatusMovedPermanently)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
