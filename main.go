package main

import (
	"DarkRoom/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", controllers.NewStatic().Home).Methods("GET")
	r.Handle("/contact", controllers.NewStatic().Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
