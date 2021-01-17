package models

import "strings"

const (
	// ErrNotFound is used when resource not found in DB
	ErrNotFound modelError = "models: resource not fould"
	// ErrIDInvalid is returned when an invalid ID is provided to a method
	// like Delete
	ErrIDInvalid modelError = "models: ID provided was invalid"
	// ErrPasswordIncorrect is returned when an invalid password is used when attempting to
	// authenticate a user
	ErrPasswordIncorrect modelError = "models: Incorrect password provided"
	// ErrEmailRequired is returned when an email address is not provided
	ErrEmailRequired modelError = "models: Email address is required"
	// ErrEmailInvalid is returned when an email address provided does not match
	// any of our requirements
	ErrEmailInvalid modelError = "models: Email address is not valid"
	// ErrEmailTaken is returned when an email address is taken
	ErrEmailTaken modelError = "models: Email address is already taken"
	// ErrPasswordTooShort is returned when an update or create is attempted with a user
	// password that is less than 8 characters
	ErrPasswordTooShort modelError = "models: Password must be at least 8 characters long"
	// ErrPasswordRequired is returned when a create is attempted without a user password provided
	ErrPasswordRequired modelError = "models: Password is required"
	// ErrRememberTooShort is returned when a remember token is not at least 32 bytes
	ErrRememberTooShort modelError = "models: Remember token must be at least 32 bytes"
	// ErrRememberRequired is returned when a create or update is attempted without a user remember token hash
	ErrRememberRequired modelError = "models: Remember token is required"
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
