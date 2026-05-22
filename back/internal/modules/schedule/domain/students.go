package schedule_domain

import "github.com/google/uuid"

var ErrstudentNotFound string = "student not found"

type Students interface {
	Get(studentId uuid.UUID) (StudentId, error)
}
