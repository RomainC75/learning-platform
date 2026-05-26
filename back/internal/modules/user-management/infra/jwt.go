package user_mngmt_infra

import (
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) CreateToken(email string, userRole user_mngmt_domain.UserRole) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_email": email,
			"user_role":  userRole,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
