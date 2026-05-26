package user_mngmt_domain

import (
	"errors"
)

var (
	ErrCreatingToken = errors.New("could not create token")
)

type Jwt interface {
	CreateToken(email string, userRole UserRole) (string, error)
}
