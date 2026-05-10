package schedule_unit

import (
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	utils_time "language-learning/utils/time"
	"time"

	"github.com/google/uuid"
)

var (
	studentUuid   = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	professorUuid = uuid.MustParse("22ce4a72-2978-4f19-9181-e59963add95b")
)

type TestDriver struct {
	professor *schedule_domain.Professor
	student   *schedule_domain.Student
}

func NewTestDriver() *TestDriver {
	professor := schedule_domain.NewProfessor(professorUuid, "big", "brother", []schedule_domain.Reservation{
		schedule_domain.NewReservation(
			time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC),
			utils_time.MustParseDuration("1h"),
		),
	})
	student := schedule_domain.NewStudent(studentUuid, "john", "Doe", []schedule_domain.Reservation{})

	return &TestDriver{
		professor: professor,
		student:   student,
	}

}

func (td *TestDriver) CreateReservationsForProfessor(rezas []schedule_domain.Reservation) *TestDriver {
	td.professor.ResetSchedule()
	for _, reza := range rezas {
		td.professor.MustAddSchedule(reza)
	}
	return td
}

func (td *TestDriver) CreateReservationsForStudent(rezas ...schedule_domain.Reservation) *TestDriver {
	td.student.ResetSchedule()
	for _, reza := range rezas {
		td.student.MustAddSchedule(reza)
	}
	return td
}

func (td *TestDriver) NewScheduleService() *schedule_application.ReservationSrv {
	professors := schedule_infra.NewInMemProfRepo(td.professor, false)
	students := schedule_infra.NewInMemStudentRepo(td.student, false)
	return schedule_application.NewReservationSrv(professors, students)

}
