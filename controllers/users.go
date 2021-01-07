package controllers

import (
	"DarkRoom/views"
	"fmt"
	"net/http"
)

// The Users structure
type Users struct {
	NewView *views.View
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
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	u.NewView.Render(w, nil)
}

// This is used to process the signup form when a user tries to create a new user account
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a message to pretend that we created the user account")
}
