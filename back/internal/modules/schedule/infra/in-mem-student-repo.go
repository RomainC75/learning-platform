package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"

	"github.com/google/uuid"
)

type InMemStudentRepo struct {
	expectedStudent *schedule_domain.Student
	isErrorExpected bool
}

func NewInMemStudentRepo(expectedStudent *schedule_domain.Student, isErrorExpected bool) *InMemStudentRepo {
	return &InMemStudentRepo{
		expectedStudent: expectedStudent,
		isErrorExpected: isErrorExpected,
	}
}

func (studentRepo *InMemStudentRepo) Get(id uuid.UUID) (*schedule_domain.Student, error) {
	if studentRepo.isErrorExpected {
		return nil, errors.New(schedule_domain.ErrstudentNotFound)
	}
	return studentRepo.expectedStudent, nil
}
