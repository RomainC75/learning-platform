package schedule_application

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
	auth_jwt "language-learning/internal/auth/jwt"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	"time"

	"github.com/google/uuid"
)

type ScheduleSrv struct {
	professors schedule_domain.Professors
}

type CreateFreeScheduleResponse struct {
	Status      bool          `json:"status"`
	StudentId   uuid.UUID     `json:"student_id"`
	ProfessorId uuid.UUID     `json:"professor_id"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
}

func NewScheduleSrv(professors schedule_domain.Professors, students schedule_domain.Students) *ScheduleSrv {
	return &ScheduleSrv{
		professors: professors,
	}
}

func (ScheduleSrv *ScheduleSrv) CreateBooking(ctx context.Context, createBookingRequest dtos.CreateBookingRequest) (CreateFreeScheduleResponse, error) {
	foundProfessor, err := ScheduleSrv.professors.Get(createBookingRequest.ProfessorId)
	if err != nil {
		return CreateFreeScheduleResponse{}, err
	}

	studentId, _ := ctx.Value(auth_jwt.UserId).(uuid.UUID)

	newBooking := schedule_domain.NewBooking(createBookingRequest.Date, createBookingRequest.Duration, studentId)
	err = foundProfessor.Schedule(newBooking)
	if err != nil {
		return CreateFreeScheduleResponse{}, err
	}

	err = ScheduleSrv.professors.AddBooking(foundProfessor, newBooking)
	if err != nil {
		return CreateFreeScheduleResponse{}, err
	}

	return CreateFreeScheduleResponse{
		Status:      true,
		StudentId:   studentId,
		ProfessorId: foundProfessor.Id(),
		Date:        createBookingRequest.Date,
		Duration:    createBookingRequest.Duration,
	}, nil
}
