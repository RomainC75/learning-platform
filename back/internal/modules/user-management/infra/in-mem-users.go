package user_mngmt_infra

import (
	"errors"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type InMemUsers struct {
	expectedUser  *user_mngmt_domain.User
	expectedError bool
}

func NewInMemUsers(expectedUser *user_mngmt_domain.User, expectSignupError bool) *InMemUsers {
	return &InMemUsers{
		expectedUser:  expectedUser,
		expectedError: expectSignupError,
	}
}

func (umu *InMemUsers) GetById(userId uuid.UUID) (*user_mngmt_domain.User, error) {
	return umu.expectedUser, nil
}

func (umu *InMemUsers) GetByEmail(userEmail string) (*user_mngmt_domain.User, error) {
	if umu.expectedError {
		return nil, errors.New(user_mngmt_domain.ErrEmailAlreadyUsed)
	}
	return umu.expectedUser, nil
}

func (umu *InMemUsers) SaveUser(newUser *user_mngmt_domain.User, newEncryptedPass string) error {
	return nil
}
