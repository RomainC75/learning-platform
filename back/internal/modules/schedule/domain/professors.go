package schedule_domain

import (
	"github.com/google/uuid"
)

var ErrProfessorNotFound string = "professor not found"

type Professors interface {
	Get(professorId uuid.UUID) (*Professor, error)
	AddReservation(professor *Professor, reservation Reservation) error
}
