package user_mngmt_infra

import (
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type InMemJWT struct {
	expectedToken    string
	expectedError    bool
	expectedJwtClaim user_mngmt_domain.JwtClaim
}

func NewInMemJWT(expectedToken string, expectedErrorOpt ...bool) *InMemJWT {
	var expectedError bool
	if len(expectedErrorOpt) > 0 {
		expectedError = expectedErrorOpt[0]
	}
	return &InMemJWT{
		expectedToken: expectedToken,
		expectedError: expectedError,
	}
}

func (imj *InMemJWT) CreateToken(userId uuid.UUID, email string, userRole user_mngmt_domain.UserRole) (string, error) {
	if imj.expectedError {
		return "", user_mngmt_domain.ErrCreatingToken
	}
	return imj.expectedToken, nil
}

func (imj *InMemJWT) GetClaimsFromToken(tokenString string) (user_mngmt_domain.JwtClaim, error) {
	return imj.expectedJwtClaim, nil
}
