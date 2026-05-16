package schedule_unit

import (
	"context"
	auth_jwt "language-learning/internal/auth/jwt"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	utils_time "language-learning/utils/time"
	"time"
)

type BookingTestDriver struct {
	professor *schedule_domain.Professor
	student   *schedule_domain.Student
}

func NewBookingTestDriver() *BookingTestDriver {
	student := schedule_domain.NewStudent(studentUuid, "john", "Doe", []schedule_domain.Booking{})
	professor := schedule_domain.NewProfessor(professorUuid, "big", "brother", schedule_domain.Schedule{}, []schedule_domain.Booking{
		schedule_domain.NewBooking(
			time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC),
			utils_time.MustParseDuration("1h"),
			otherStudentUuid,
		),
	})

	return &BookingTestDriver{
		professor: professor,
		student:   student,
	}

}

func (td *BookingTestDriver) CreateBookingsForProfessor(rezas []schedule_domain.Booking) *BookingTestDriver {
	td.professor.ResetBooking()
	for _, reza := range rezas {
		td.professor.MustAddBooking(reza)
	}
	return td
}

func (td *BookingTestDriver) CreateBookingsForStudent(rezas ...schedule_domain.Booking) *BookingTestDriver {
	td.student.ResetBooking()
	for _, reza := range rezas {
		td.student.MustAddBooking(reza)
	}
	return td
}

func (td *BookingTestDriver) NewScheduleService() *schedule_application.BookingSrv {
	professors := schedule_infra.NewInMemProfRepo(td.professor, false)
	students := schedule_infra.NewInMemStudentRepo(td.student, false)
	return schedule_application.NewBookingSrv(professors, students)
}

func (td *BookingTestDriver) BuildStudentContext() context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, auth_jwt.UserId, studentUuid)
}
