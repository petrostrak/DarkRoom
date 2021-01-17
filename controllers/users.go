package controllers

import (
	"DarkRoom/models"
	"DarkRoom/rand"
	"DarkRoom/views"
	"fmt"
	"net/http"
)

// The Users structure
type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
}

// The SignupForm struct for signing up
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

// New is used to render the form where a user can create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type Alert struct {
		Level   string
		Message string
	}
	type Data struct {
		Alert Alert
		Yield interface{}
	}
	a := Alert{
		Level:   "success",
		Message: "Successfully render a dynamic alert!",
	}
	d := Data{
		Alert: a,
		Yield: "Hello!",
	}
	if err := u.NewView.Render(w, d); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user tries to create a new user account
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := u.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// LoginForm struct
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login is used to verify the provided email address and password and then
// log the user in if they are correct
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrPasswordIncorrect:
			fmt.Fprintln(w, "Invalid password provided.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// if no error occured during the login, then create the cookie
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// create cookie method
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// CookieTest is used to display cookies set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}
