package schedule_application

import (
	"context"
	dtos_requests "language-learning/internal/api/dtos/request"
	"language-learning/internal/api/middlewares"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	shared_domain_uuid "language-learning/internal/modules/shared/domain/uuid"
	"time"

	"github.com/google/uuid"
)

type BookingSrv struct {
	idGenerator   shared_domain_uuid.UuidGenerator
	userReader    schedule_domain.UserReader
	professors    schedule_domain.Professors
	students      schedule_domain.Students
	bookings      schedule_domain.Bookings
	timeGenerator shared_domain_time.TimeGenerator
}

type CreateBookingResponse struct {
	Status      bool          `json:"status"`
	StudentId   uuid.UUID     `json:"student_id"`
	ProfessorId uuid.UUID     `json:"professor_id"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
}

func NewBookingSrv(uuidGenerator shared_domain_uuid.UuidGenerator, timeGenerator shared_domain_time.TimeGenerator, professors schedule_domain.Professors, students schedule_domain.Students, bookings schedule_domain.Bookings) *BookingSrv {
	return &BookingSrv{
		idGenerator:   uuidGenerator,
		timeGenerator: timeGenerator,
		professors:    professors,
		students:      students,
		bookings:      bookings,
	}
}

func (bookingSrv *BookingSrv) CreateBooking(ctx context.Context, createBookingRequest dtos_requests.CreateBookingRequest) (CreateBookingResponse, error) {
	foundProfessor, err := bookingSrv.professors.Get(createBookingRequest.ProfessorId)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	studentuuid, _ := ctx.Value(middlewares.UserIdKey).(uuid.UUID)
	studentId, err := bookingSrv.students.Get(studentuuid)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	newBookingUuid := bookingSrv.idGenerator.Generate()
	dtr := shared_domain_time.NewDateTimeRange(createBookingRequest.Date, createBookingRequest.Duration)
	newBooking, err := schedule_domain.NewBooking(newBookingUuid, dtr, foundProfessor.Id(), studentId, bookingSrv.timeGenerator.Now())
	if err != nil {
		return CreateBookingResponse{}, err
	}

	bookings := bookingSrv.bookings.GetBookings(foundProfessor)
	err = foundProfessor.IsBookable(bookings, newBooking)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	err = bookingSrv.bookings.SetBooking(newBooking)
	if err != nil {
		return CreateBookingResponse{}, err
	}

	return CreateBookingResponse{
		Status:      true,
		StudentId:   uuid.UUID(studentId),
		ProfessorId: uuid.UUID(foundProfessor.Id()),
		Date:        createBookingRequest.Date,
		Duration:    createBookingRequest.Duration,
	}, nil
}
