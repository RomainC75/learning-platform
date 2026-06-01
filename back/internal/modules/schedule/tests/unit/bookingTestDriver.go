package schedule_unit

import (
	"context"
	dtos_requests "language-learning/internal/api/dtos/request"
	"language-learning/internal/api/middlewares"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	shared_infra "language-learning/internal/modules/shared/infra"
)

type BookingTestDriver struct {
	uuidGenerator *shared_infra.InMemUuidGenerator
	professor     *schedule_domain.Professor
	student       schedule_domain.StudentId
	bookings      *schedule_infra.InMemBookingRepo
	bookingSrv    *schedule_application.BookingSrv
}

func NewBookingTestDriver() *BookingTestDriver {
	uuidGenerator := shared_infra.NewInMemUuidGenerator(newBookingUuid)

	bookings := schedule_infra.NewInMemBookingRepo()

	return &BookingTestDriver{
		uuidGenerator: uuidGenerator,
		professor:     professor,
		student:       student,
		bookings:      bookings,
	}

}

func (td *BookingTestDriver) CreateBookingsForProfessor(rezas []schedule_domain.Booking) *BookingTestDriver {
	td.bookings.ResetBooking()
	for _, reza := range rezas {
		td.bookings.MustAddBooking(reza)
	}
	return td
}

func (td *BookingTestDriver) NewScheduleService() *schedule_application.BookingSrv {
	professors := schedule_infra.NewInMemProfRepo(td.professor, false)
	students := schedule_infra.NewInMemStudentRepo(td.student, false)
	timeGenerator := shared_infra.NewDeterministicTimeGenerator(nowMarch1)
	td.bookingSrv = schedule_application.NewBookingSrv(td.uuidGenerator, timeGenerator, professors, students, td.bookings)
	return td.bookingSrv
}

func (td *BookingTestDriver) BuildStudentContext() context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, middlewares.UserIdKey, studentUuid)
}

func (td *BookingTestDriver) SavedBookings() []schedule_domain.Booking {
	return td.bookings.ListBookings()
}

func (td *BookingTestDriver) CreateBooking(ctx context.Context, createBookingRequest dtos_requests.CreateBookingRequest) (schedule_application.CreateBookingResponse, error) {
	return td.bookingSrv.CreateBooking(ctx, createBookingRequest)
}
