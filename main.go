package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if (r.URL.Path) == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if (r.URL.Path) == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@darkroom.com\">support@darkroom.com</a>.")
	} else {
		// status codes should be written before any template rendering
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1>")
	}
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", myHandler)
	// http.HandleFunc("/", myHandler)
	http.ListenAndServe(":3000", mux)
}
