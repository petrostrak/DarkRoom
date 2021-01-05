package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
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

func secondHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is a practice web page, cheers!</h1>")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s\n", ps.ByName("name"))
}

// go get github.com/julienschmidt/httprouter
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", myHandler)
	// http.HandleFunc("/", myHandler)
	http.ListenAndServe(":3000", r)
}
