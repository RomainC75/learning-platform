package schedule_domain

import (
	"github.com/google/uuid"
)

var ErrProfessorNotFound string = "professor not found"

type Professors interface {
	Get(professorId uuid.UUID) (*Professor, error)
	Save(professor *Professor) error
}
