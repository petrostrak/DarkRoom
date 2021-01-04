package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":3000", nil)
}
