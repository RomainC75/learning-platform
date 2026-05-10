package schedule_domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrProfessorNotSchedulable = "professor not schedulable"

type Professor struct {
	id           uuid.UUID
	firstName    string
	lastName     string
	reservations []Reservation
}

func NewProfessor(id uuid.UUID, firstName string, lastName string, reservation []Reservation) *Professor {
	return &Professor{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		reservations: reservation,
	}
}

func (professor *Professor) Schedule(newReservation *Reservation) error {
	isSchedulable := newReservation.IsReservableIn(professor.reservations)
	if !isSchedulable {
		return errors.New(ErrProfessorNotSchedulable)
	}
	professor.reservations = append(professor.reservations, *newReservation)
	return nil
}
