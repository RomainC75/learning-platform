package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"

	"github.com/google/uuid"
)

type FakeStudentRepo struct {
	expectedStudent schedule_domain.StudentId
	isErrorExpected bool
}

func NewFakeStudentRepo(expectedStudent schedule_domain.StudentId, isErrorExpected bool) *FakeStudentRepo {
	return &FakeStudentRepo{
		expectedStudent: expectedStudent,
		isErrorExpected: isErrorExpected,
	}
}

func (studentRepo *FakeStudentRepo) Get(id uuid.UUID) (schedule_domain.StudentId, error) {
	if studentRepo.isErrorExpected {
		return schedule_domain.StudentId{}, errors.New(schedule_domain.ErrstudentNotFound)
	}
	return studentRepo.expectedStudent, nil
}
