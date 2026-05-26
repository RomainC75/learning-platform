package user_mngmt_domain

import "errors"

var (
	ErrTryingToGenerateBcryptPassword = errors.New("bcrypt : error trying to generate password. maximum 72 bytes long")
	ErrInvalidHashedPass              = errors.New("hashedPasword invalid")
)

type Bcrypt interface {
	CompareHashAndPassword(hashedPassword, password []byte) error
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
}
