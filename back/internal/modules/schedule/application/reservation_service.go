package schedule_application

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
	schedule_domain "language-learning/internal/modules/schedule/domain"
)

type ReservationSrv struct {
	professors schedule_domain.Professors
	students   schedule_domain.Students
}

func NewReservationSrv(professors schedule_domain.Professors, students schedule_domain.Students) *ReservationSrv {
	return &ReservationSrv{
		professors: professors,
		students:   students,
	}
}

func (reservationSrv *ReservationSrv) CreateReservation(ctx context.Context, createReservationRequest dtos.CreateReservationRequest) (string, error) {
	foundProfessor, err := reservationSrv.professors.Get(createReservationRequest.ProfessorId)
	if err != nil {
		return "", err
	}

	newReservation := schedule_domain.NewReservation(createReservationRequest.Date, createReservationRequest.Duration)
	err = foundProfessor.Schedule(newReservation)

	if err != nil {
		return "", err
	}

	// fmt.Fprintf("stu", )
	return "scheduled", nil
}
