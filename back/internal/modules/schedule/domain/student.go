package schedule_domain

import (
	"github.com/google/uuid"
)

type Student struct {
	id        uuid.UUID
	firstName string
	lastName  string
}

func NewStudent(id uuid.UUID, firstName string, lastName string) *Student {
	return &Student{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
	}
}

// == Snapshot ==

type StudentSnapshot struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}

func (s Student) ToSnapshot() StudentSnapshot {
	return StudentSnapshot{
		Id:        s.id,
		FirstName: s.firstName,
		LastName:  s.lastName,
	}
}
