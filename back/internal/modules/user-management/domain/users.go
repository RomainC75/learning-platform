package user_mngmt_domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmailAlreadyUsed       = errors.New("email already used")
	ErrWrongEmailOrPassword   = errors.New("wrong email or password")
	ErrTryingtoSaveTheNewUser = errors.New("error trying to save the new user")
	ErrUserNotFound           = errors.New("error user not found")
)

type Users interface {
	GetById(userId uuid.UUID) (*User, error)
	GetByEmail(userEmail string) (*User, error)
	Save(newUser *User) error
}
