package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"

	"github.com/google/uuid"
)

type FakeProfRepo struct {
	expectedProfessor *schedule_domain.Professor
	addedBooking      []schedule_domain.Booking
	savedSchedule     schedule_domain.Schedule
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

func (profRepo *FakeProfRepo) AddBooking(professor *schedule_domain.Professor, booking schedule_domain.Booking) error {
	return nil
}

func (profRepo *FakeProfRepo) ReplaceSchedule(professor *schedule_domain.Professor, schedule schedule_domain.Schedule) error {
	profRepo.savedSchedule = schedule
	return nil
}

func (profRepo *FakeProfRepo) GetSavedSchedule() schedule_domain.Schedule {
	return profRepo.savedSchedule
}
