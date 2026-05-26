package user_mngmt_domain

import (
	"github.com/google/uuid"
)

type UserRole int

const (
	IsProfessor UserRole = iota
	IsStudent
)

type User struct {
	id        uuid.UUID
	email     string
	password  string
	firstName string
	lastName  string
	userRole  UserRole
}

func NewUser(id uuid.UUID, email string, password string, firstName string, lastName string, userRole UserRole) *User {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		firstName: firstName,
		lastName:  lastName,
		userRole:  userRole,
	}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) IsProfessor() bool {
	return u.userRole == IsProfessor
}

func (u *User) IsStudent() bool {
	return u.userRole == IsStudent
}

func (u *User) IsPasswordValid(compFn func(hashedPassword []byte, password []byte) error, clearPass string) error {
	return compFn([]byte(u.password), []byte(clearPass))
}
