package user_mngmt_infra

import (
	"fmt"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type InMemUsers struct {
	expectedUser          *user_mngmt_domain.User
	savedPassword         string
	expectedGetError      bool
	expectedSaveUserError bool
}

func NewInMemUsers(expectedUser *user_mngmt_domain.User, expectGetError bool, expectedSaveUserError bool) *InMemUsers {
	return &InMemUsers{
		expectedUser:          expectedUser,
		expectedGetError:      expectGetError,
		expectedSaveUserError: expectedSaveUserError,
	}
}

func (umu *InMemUsers) GetById(userId uuid.UUID) (*user_mngmt_domain.User, error) {
	return umu.expectedUser, nil
}

func (umu *InMemUsers) GetByEmail(userEmail string) (*user_mngmt_domain.User, error) {
	if umu.expectedGetError {
		return nil, user_mngmt_domain.ErrUserNotFound
	}
	return umu.expectedUser, nil
}

func (umu *InMemUsers) Save(newUser *user_mngmt_domain.User) error {
	if umu.expectedSaveUserError {
		return user_mngmt_domain.ErrTryingtoSaveTheNewUser
	}
	fmt.Println("SAVE : ", newUser)
	umu.expectedUser = newUser
	return nil
}

func (umu *InMemUsers) GetLastSavedUser() *user_mngmt_domain.User {
	fmt.Println("SAVD : ", umu.expectedUser)
	return umu.expectedUser
}
