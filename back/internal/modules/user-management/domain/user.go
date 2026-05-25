package user_mngmt_domain

import "github.com/google/uuid"

type User struct {
	id          uuid.UUID
	email       string
	password    string
	firstName   string
	lastName    string
	isProfessor bool
}

func NewUser(id uuid.UUID, email string, password string, firstName string, lastName string, isProfessor bool) *User {
	return &User{
		id:          id,
		email:       email,
		password:    password,
		firstName:   firstName,
		lastName:    lastName,
		isProfessor: isProfessor,
	}
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) IsProfessor() bool {
	return u.isProfessor
}

// func (u *User) IsPasswordValid(compareFn ) bool {

// }
