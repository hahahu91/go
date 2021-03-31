package main

import (
	"net/http"
)

const (
	DefaultFallbackUrl = "https://golang.org/"
)

func main () {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	shortUrl := map[string]string {
		"/go-http":    "https://golang.org/pkg/net/http/",
		"/go-gophers": "https://github.com/shalakhin/gophericons/blob/master/preview.jpg",
	}
	for k, v := range shortUrl {
		if r.URL.Path == k {
			http.Redirect(w, r, v, http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, DefaultFallbackUrl, http.StatusFound)
	return
}