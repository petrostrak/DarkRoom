package controllers

import (
	"DarkRoom/views"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

// The Users structure
type Users struct {
	NewView *views.View
}

// The SignupForm struct for signing up
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//This is used to render the form where a user can create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	u.NewView.Render(w, nil)
}

// This is used to process the signup form when a user tries to create a new user account
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	decoder := schema.NewDecoder()
	var form SignupForm

	if err := decoder.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
