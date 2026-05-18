package schedule_domain

import (
	"github.com/google/uuid"
)

var ErrProfessorNotSchedulable = "professor not schedulable"

type Professor struct {
	id        uuid.UUID
	firstName string
	lastName  string
	schedule  Schedule
}

func NewProfessor(id uuid.UUID, firstName string, lastName string, schedule Schedule) *Professor {
	return &Professor{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		schedule:  schedule,
	}
}

func (professor *Professor) SetSchedule(schedule Schedule) {
	professor.schedule = schedule
}

func (professor *Professor) Id() uuid.UUID {
	return professor.id
}
