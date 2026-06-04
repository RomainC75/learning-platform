package user_mngmt_infra

import (
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type InMemUsers struct {
	users []*user_mngmt_domain.User
}

type InMemUsersErrors struct {
	Users []*user_mngmt_domain.User
}

func NewInMemUsers() *InMemUsers {
	return &InMemUsers{}
}

func (umu *InMemUsers) GetById(userId uuid.UUID) (*user_mngmt_domain.User, error) {
	for _, u := range umu.users {
		if userId == u.Id() {
			return u, nil
		}
	}
	return nil, user_mngmt_domain.ErrUserNotFound
}

func (umu *InMemUsers) GetByEmail(userEmail string) (*user_mngmt_domain.User, error) {
	for _, u := range umu.users {
		if userEmail == u.Email() {
			return u, nil
		}
	}
	return nil, user_mngmt_domain.ErrUserNotFound
}

func (umu *InMemUsers) Save(newUser *user_mngmt_domain.User) error {
	if _, err := umu.GetByEmail(newUser.Email()); err == nil {
		return user_mngmt_domain.ErrEmailAlreadyUsed
	}
	umu.users = append(umu.users, newUser)
	return nil
}
