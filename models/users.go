package models

import (
	"DarkRoom/hash"
	"DarkRoom/rand"
	"errors"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is used when resource not found in DB
	ErrNotFound = errors.New("models: resource not fould")
	// ErrInvalidID is returned when an invalid ID is provided to a method
	// like Delete
	ErrInvalidID = errors.New("models: ID provided was invalid")
	// ErrInvalidPassword is returned when an invalid password is used when attempting to
	// authenticate a user
	ErrInvalidPassword = errors.New("models: Incorrect password provided")
)

const userPwdPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

// User represents the user model stored in our DB
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unigue"`
	// with "-" we say gorm not to save thi s field in the DB
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique"`
}

// UserDB is used to interact with the users DB
//
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// For single user quieries, any error but ErrNotFound should probably
// result in a 500 error.
type UserDB interface {
	// Methods for queryng for single users
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// Used to close DB connection
	Close() error

	// Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

// UserService is a set of methods used to manipulate and
// work with the user model
type UserService interface {
	// Authenticate wii verify the provided email address and
	// password are correct. If the are correct, the user 
	// corresponding to that email will be returned 
	Authenticate(email, password string) (*User, error)
	UserDB
}

// NewUserService instantiates a new User service
func NewUserService(connectionInfo string) (UserService, error) {
	ug, err :=newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	return &userService{
		UserDB: &userValidator{
			UserDB: ug,
		},
	}, nil
}

// DestructiveReset drops user table and rebuilds it
func (ug *userGorm) DestructiveReset() error {
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return ug.AutoMigrate()
}

// AutoMigrate will migrate our data
func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

// UserService struct
type userService struct {
	UserDB
}

// Authenticate can be used to authenticate the user with the provided email address and password
func (us *userService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwdPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}
	return foundUser, nil
}

var _ UserDB = &userValidator{}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &userGorm{
		db: db,
		hmac: hmac,
	}, nil
}

type userValidator struct {
	UserDB
}

var _ UserDB = &userGorm{}

type userGorm struct {
	db   *gorm.DB
	hmac hash.HMAC
}

// ByID will look up a user by the id provided
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// As a general rule, any error but ErrNotFound should probably
// result in a 500 error.
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will look up a user by the email provided
// 1 - user, nil. If user is found, return a nil error
// 2 - nil, ErrNotFound. If user not found, return ErrNotFound
// 3 - otherError. If other error occurs, return error in detail
//
// As a general rule, any error but ErrNotFound should probably
// result in a 500 error.
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// ByRemember looks up a user with the given remember token and returns that user
// This method will handle hashing the token for us
func (ug *userGorm) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := ug.hmac.Hash(token)
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create doesn't return the user, instead we update the one we pass in
// therefore we use a pointer to User
func (ug *userGorm) Create(user *User) error {
	// implementing pepper
	pwBytes := []byte(user.Password + userPwdPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// convert bytes to string
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = ug.hmac.Hash(user.Remember)
	return ug.db.Create(user).Error
}

// Update will update the provided user with all of the data
// in the provided user object
func (ug *userGorm) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = ug.hmac.Hash(user.Remember)
	}
	return ug.db.Save(user).Error
}

// Delete will delete the user with the provided id
func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

// Close closes the UserService db connection
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

// first will query using the provided gorm.DB and it will get
// the first item returned and place it into dst. If nothing
// is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}