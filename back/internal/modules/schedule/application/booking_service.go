package schedule_application

import (
	"context"
	dtos_requests "language-learning/internal/api/dtos/request"
	auth_jwt "language-learning/internal/auth/jwt"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	"time"

	"github.com/google/uuid"
)

type BookingSrv struct {
	professors schedule_domain.Professors
	students   schedule_domain.Students
}

type CreateBookingResponse struct {
	Status      bool          `json:"status"`
	StudentId   uuid.UUID     `json:"student_id"`
	ProfessorId uuid.UUID     `json:"professor_id"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
}

func NewBookingSrv(professors schedule_domain.Professors, students schedule_domain.Students) *BookingSrv {
	return &BookingSrv{
		professors: professors,
		students:   students,
	}
}

func (bookingSrv *BookingSrv) CreateBooking(ctx context.Context, createBookingRequest dtos_requests.CreateBookingRequest) (CreateBookingResponse, error) {
	foundProfessor, err := bookingSrv.professors.Get(createBookingRequest.ProfessorId)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	studentId, _ := ctx.Value(auth_jwt.UserId).(uuid.UUID)

	newBooking := schedule_domain.NewBooking(createBookingRequest.Date, createBookingRequest.Duration, studentId)
	err = foundProfessor.Booking(newBooking)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	err = bookingSrv.professors.AddBooking(foundProfessor, newBooking)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	return CreateBookingResponse{
		Status:      true,
		StudentId:   studentId,
		ProfessorId: foundProfessor.Id(),
		Date:        createBookingRequest.Date,
		Duration:    createBookingRequest.Duration,
	}, nil
}
