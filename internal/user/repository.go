package user

import (
	"errors"

	"gorm.io/gorm"
)

// keep the core things here like database interaction with user
// dont write any http related code here

var ErrorAlreadyExist = errors.New("user with this email already exists")

type Repository interface {
	// create a user
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// create a user
func (r repository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {

		// handle error via error package not string
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {

			return ErrorAlreadyExist
		}

		return result.Error
	}
	return nil
}

// login
func (r repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
