package schedule_domain

import "github.com/google/uuid"

type Student struct {
	id           uuid.UUID
	firstName    string
	lastName     string
	reservations []Reservation
}

func NewStudent(id uuid.UUID, firstName string, lastName string, reservation []Reservation) *Student {
	return &Student{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		reservations: reservation,
	}
}

func (student *Student) MustAddSchedule(newReza Reservation) {
	student.reservations = append(student.reservations, newReza)
}

func (student *Student) ResetSchedule() {
	student.reservations = []Reservation{}
}
