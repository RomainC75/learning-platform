package user_mngmt_domain

import "github.com/google/uuid"

var (
	ErrEmailAlreadyUsed       = "email already used"
	ErrWrongEmailOrPassword   = "wrong email or password"
	ErrTryingtoSaveTheNewUser = "error trying to save the new user"
)

type Users interface {
	GetById(userId uuid.UUID) (*User, error)
	GetByEmail(userEmail string) (*User, error)
	SaveUser(newUser *User, newEncryptedPass string) error
}
