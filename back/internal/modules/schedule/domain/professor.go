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
	Bookings  []Booking
}

func NewProfessor(id uuid.UUID, firstName string, lastName string, Booking []Booking) *Professor {
	return &Professor{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		Bookings:  Booking,
	}
}

func (professor *Professor) Schedule(newBooking Booking) error {
	isSchedulable := newBooking.isBookableIn(professor.Bookings)
	if !isSchedulable {
		return errors.New(ErrProfessorNotSchedulable)
	}
	professor.Bookings = append(professor.Bookings, newBooking)
	return nil
}

func (professor *Professor) MustAddSchedule(newReza Booking) {
	professor.Bookings = append(professor.Bookings, newReza)
}

func (professor *Professor) ResetSchedule() {
	professor.Bookings = []Booking{}
}

func (professor *Professor) Id() uuid.UUID {
	return professor.id
}
