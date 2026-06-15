package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"

	"github.com/google/uuid"
)

type FakeProfRepo struct {
	expectedProfessor *schedule_domain.Professor
	isErrorExpected   bool
}

func NewFakeProfRepo(expectedProfessor *schedule_domain.Professor, isErrorExpected bool) *FakeProfRepo {
	return &FakeProfRepo{
		expectedProfessor: expectedProfessor,
		isErrorExpected:   isErrorExpected,
	}
}

func (profRepo *FakeProfRepo) Get(id uuid.UUID) (*schedule_domain.Professor, error) {
	if profRepo.isErrorExpected {
		return nil, errors.New(schedule_domain.ErrProfessorNotFound)
	}
	return profRepo.expectedProfessor, nil
}

func (profRepo *FakeProfRepo) GetSavedSchedule() schedule_domain.Schedule {
	return schedule_domain.NewScheduleFromSnapshot(profRepo.expectedProfessor.ToSnapshot().Schedule)
}

func (profRepo *FakeProfRepo) Save(professor *schedule_domain.Professor) error {
	profRepo.expectedProfessor = professor
	return nil
}
