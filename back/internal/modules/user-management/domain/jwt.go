package user_mngmt_domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCreatingToken = errors.New("could not create token")
)

type JwtClaim struct {
	UserId    uuid.UUID
	UserEmail string
	UserRole  UserRole
}

type Jwt interface {
	CreateToken(userId uuid.UUID, email string, userRole UserRole) (string, error)
	GetClaimsFromToken(tokenString string) (JwtClaim, error)
}
