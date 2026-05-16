package schedule_domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrProfessorNotSchedulable = "professor not schedulable"

type Professor struct {
	id        uuid.UUID
	firstName string
	lastName  string
	schedule  Schedule
	bookings  []Booking
}

func NewProfessor(id uuid.UUID, firstName string, lastName string, schedule Schedule, Booking []Booking) *Professor {
	return &Professor{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		schedule:  schedule,
		bookings:  Booking,
	}
}

func (professor *Professor) SetSchedule(schedule Schedule) {
	professor.schedule = schedule
}

func (professor *Professor) Booking(newBooking Booking) error {
	isSchedulable := newBooking.isBookableIn(professor.bookings)
	if !isSchedulable {
		return errors.New(ErrProfessorNotSchedulable)
	}
	professor.bookings = append(professor.bookings, newBooking)
	return nil
}

func (professor *Professor) MustAddBooking(newReza Booking) {
	professor.bookings = append(professor.bookings, newReza)
}

func (professor *Professor) ResetBooking() {
	professor.bookings = []Booking{}
}

func (professor *Professor) Id() uuid.UUID {
	return professor.id
}
