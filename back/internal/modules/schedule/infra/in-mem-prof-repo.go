package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"

	"github.com/google/uuid"
)

type InMemProfRepo struct {
	expectedProfessor *schedule_domain.Professor
	addedReservation  []schedule_domain.Reservation
	isErrorExpected   bool
}

func NewInMemProfRepo(expectedProfessor *schedule_domain.Professor, isErrorExpected bool) *InMemProfRepo {
	return &InMemProfRepo{
		expectedProfessor: expectedProfessor,
		isErrorExpected:   isErrorExpected,
	}
}

func (profRepo *InMemProfRepo) Get(id uuid.UUID) (*schedule_domain.Professor, error) {
	if profRepo.isErrorExpected {
		return nil, errors.New(schedule_domain.ErrProfessorNotFound)
	}
	return profRepo.expectedProfessor, nil
}

func (profRepo *InMemProfRepo) AddReservation(professor *schedule_domain.Professor, reservation schedule_domain.Reservation) error {
	return nil
}
