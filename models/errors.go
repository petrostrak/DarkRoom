package models

import "strings"

const (
	ErrNotFound          modelError   = "models: resource not fould"
	ErrPasswordIncorrect modelError   = "models: Incorrect password provided"
	ErrEmailRequired     modelError   = "models: Email address is required"
	ErrEmailInvalid      modelError   = "models: Email address is not valid"
	ErrEmailTaken        modelError   = "models: Email address is already taken"
	ErrPasswordTooShort  modelError   = "models: Password must be at least 8 characters long"
	ErrPasswordRequired  modelError   = "models: Password is required"
	ErrRememberTooShort  privateError = "models: Remember token must be at least 32 bytes"
	ErrRememberRequired  privateError = "models: Remember token is required"
	ErrIDInvalid         privateError = "models: ID provided was invalid"
	ErrUserIDRequired    privateError = "models: User ID is required"
	ErrTitleRequired     modelError   = "models: Title is requiered"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
