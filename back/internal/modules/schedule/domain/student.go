package schedule_domain

import "github.com/google/uuid"

type Student struct {
	id        uuid.UUID
	firstName string
	lastName  string
	Bookings  []Booking
}

func NewStudent(id uuid.UUID, firstName string, lastName string, Booking []Booking) *Student {
	return &Student{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		Bookings:  Booking,
	}
}

func (student *Student) MustAddSchedule(newReza Booking) {
	student.Bookings = append(student.Bookings, newReza)
}

func (student *Student) ResetSchedule() {
	student.Bookings = []Booking{}
}
