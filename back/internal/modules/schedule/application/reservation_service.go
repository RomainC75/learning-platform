package schedule_application

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
	auth_jwt "language-learning/internal/auth/jwt"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	"time"

	"github.com/google/uuid"
)

type ReservationSrv struct {
	professors schedule_domain.Professors
	students   schedule_domain.Students
}

type CreateReservationResponse struct {
	Status      bool          `json:"status"`
	StudentId   uuid.UUID     `json:"student_id"`
	ProfessorId uuid.UUID     `json:"professor_id"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
}

func NewReservationSrv(professors schedule_domain.Professors, students schedule_domain.Students) *ReservationSrv {
	return &ReservationSrv{
		professors: professors,
		students:   students,
	}
}

func (reservationSrv *ReservationSrv) CreateReservation(ctx context.Context, createReservationRequest dtos.CreateReservationRequest) (CreateReservationResponse, error) {
	foundProfessor, err := reservationSrv.professors.Get(createReservationRequest.ProfessorId)
	if err != nil {
		return CreateReservationResponse{}, err
	}

	studentId, _ := ctx.Value(auth_jwt.UserId).(uuid.UUID)

	newReservation := schedule_domain.NewReservation(createReservationRequest.Date, createReservationRequest.Duration, studentId)
	err = foundProfessor.Schedule(newReservation)
	if err != nil {
		return CreateReservationResponse{}, err
	}

	err = reservationSrv.professors.AddReservation(foundProfessor, newReservation)
	if err != nil {
		return CreateReservationResponse{}, err
	}

	return CreateReservationResponse{
		Status:      true,
		StudentId:   studentId,
		ProfessorId: foundProfessor.Id(),
		Date:        createReservationRequest.Date,
		Duration:    createReservationRequest.Duration,
	}, nil
}
