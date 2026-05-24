package schedule_domain

import "github.com/google/uuid"

type UserView struct {
	Id          uuid.UUID
	Email       string
	Password    string
	FirstName   string
	LastName    string
	IsProfessor bool
}

type UserReader interface {
	Get(userId uuid.UUID) (UserView, error)
}
