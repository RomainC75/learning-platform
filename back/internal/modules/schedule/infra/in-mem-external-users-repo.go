package schedule_infra

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

type ExternalUsers struct {
	users user_mngmt_domain.Users
}

func NewExternalUsers(users user_mngmt_domain.Users) *ExternalUsers {
	return &ExternalUsers{
		users: users,
	}
}

func (eu *ExternalUsers) Get(userId uuid.UUID) (schedule_domain.UserView, error) {
	foundUser, err := eu.users.GetById(userId)
	if err != nil {
		return schedule_domain.UserView{}, err
	}
	return schedule_domain.UserView{
		Id:        userId,
		FirstName: foundUser.FirstName(),
		LastName:  foundUser.LastName(),
	}, nil
}
