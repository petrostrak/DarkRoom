package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// when resource not found in DB
	ErrNotFound = errors.New("models: resource not fould")
	// is returned when an invalid ID is provided to a method
	// like Delete
	ErrInvalidID = error.New("models: ID provided was invalid")
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil
}

// drops user table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

type UserService struct {
	db *gorm.DB
}

// ById will look up a user by the id provided
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// As a general rule, any error but ErrNotFound should probably
// result in a 500 error.
func (us UserService) ById(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// will look up a user by the email provided
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// As a general rule, any error but ErrNotFound should probably
// result in a 500 error.
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// will query using the provided gorm.DB and it will get
// the first item returned and place it into dst. If nothing
// is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// we don't return the user, instead we update the one we pass in
// therefore we use a pointer to User
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// will update the provided user with all of the data
// in the provided user object
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// will delete the user with the provided id
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// closes the UserService db connection
func (us *UserService) Close() error {
	return us.db.Close()
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unigue"`
}
