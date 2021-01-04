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
	}
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":3000", nil)
}
