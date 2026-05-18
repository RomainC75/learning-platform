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

// func (student *Student) MustAddBooking(newReza Booking) {
// 	student.Bookings = append(student.Bookings, newReza)
// }

// func (student *Student) ResetBooking() {
// 	student.Bookings = []Booking{}
// }
