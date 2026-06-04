package user_mngmt_infra

import (
	"fmt"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type FakeUsers struct {
	expectedUser  *user_mngmt_domain.User
	savedPassword string
	FakeUsersErrors
}

type FakeUsersErrors struct {
	ExpectedGetError      bool
	ExpectedSaveUserError bool
}

func NewFakeUsers(fakeUsersErrors FakeUsersErrors) *FakeUsers {
	return &FakeUsers{
		FakeUsersErrors: fakeUsersErrors,
	}
}

func (umu *FakeUsers) GetById(userId uuid.UUID) (*user_mngmt_domain.User, error) {
	return umu.expectedUser, nil
}

func (umu *FakeUsers) GetByEmail(userEmail string) (*user_mngmt_domain.User, error) {
	if umu.ExpectedGetError {
		return nil, user_mngmt_domain.ErrUserNotFound
	}
	return umu.expectedUser, nil
}

func (umu *FakeUsers) Save(newUser *user_mngmt_domain.User) error {
	if umu.ExpectedSaveUserError {
		return user_mngmt_domain.ErrTryingtoSaveTheNewUser
	}
	umu.expectedUser = newUser
	return nil
}

func (umu *FakeUsers) GetLastSavedUser() *user_mngmt_domain.User {
	fmt.Println("SAVED : ", umu.expectedUser)
	return umu.expectedUser
}
