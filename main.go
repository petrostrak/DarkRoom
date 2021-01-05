package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@darkroom.com\">support@darkroom.com</a>.")
}

// func error(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1>")
// }

func secondHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is a practice web page, cheers!</h1>")
}

// go get github.com/julienschmidt/httprouter
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	// http.HandleFunc("/", myHandler)
	http.ListenAndServe(":3000", r)
}
