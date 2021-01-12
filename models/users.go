package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// when resource not found in DB
	ErrNotFound = errors.New("models: resource not fould")
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

// ById will look up by the id provided
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// As a general rule, any error but ErrNotFound should probably
// result in a 500 error.
func (us UserService) ById(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
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
