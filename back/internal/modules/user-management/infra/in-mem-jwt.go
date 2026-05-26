package user_mngmt_infra

import user_mngmt_domain "language-learning/internal/modules/user-management/domain"

type InMemJWT struct {
	expectedToken string
	expectedError bool
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

func (imj *InMemJWT) CreateToken(username string) (string, error) {
	if imj.expectedError {
		return "", user_mngmt_domain.ErrCreatingToken
	}
	return imj.expectedToken, nil
}
