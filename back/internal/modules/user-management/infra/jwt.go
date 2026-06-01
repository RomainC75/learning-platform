package user_mngmt_infra

import (
	"errors"
	"fmt"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte("secret-key")

type Claim struct {
	jwt.RegisteredClaims
	ID    uuid.UUID
	Email string
	Role  user_mngmt_domain.UserRole
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) CreateToken(userId uuid.UUID, email string, userRole user_mngmt_domain.UserRole) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               userId,
		Email:            email,
		Role:             userRole,
	})

	signedString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func (j *JWT) GetClaimsFromToken(tokenString string) (user_mngmt_domain.JwtClaim, error) {
	var userClaim Claim

	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return user_mngmt_domain.JwtClaim{}, err
	}

	if !token.Valid {
		return user_mngmt_domain.JwtClaim{}, errors.New("invalid token")
	}

	return user_mngmt_domain.JwtClaim{
		UserId:    userClaim.ID,
		UserEmail: userClaim.Email,
		UserRole:  userClaim.Role,
	}, nil
}
