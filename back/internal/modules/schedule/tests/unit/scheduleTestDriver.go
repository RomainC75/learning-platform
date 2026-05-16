package schedule_unit

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
)

type ScheduleTestDriver struct {
	professor *schedule_domain.Professor
	student   *schedule_domain.Student
}

// func NewScheduleTestDriver() *ScheduleTestDriver {
// 	student := schedule_domain.NewStudent(studentUuid, "john", "Doe", []schedule_domain.Booking{})
// 	professor := schedule_domain.NewProfessor(professorUuid, "big", "brother", []schedule_domain.Booking{
// 		schedule_domain.NewBooking(
// 			time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC),
// 			utils_time.MustParseDuration("1h"),
// 			otherStudentUuid,
// 		),
// 	})

// 	return &ScheduleTestDriver{
// 		professor: professor,
// 		student:   student,
// 	}

// }

// func (td *ScheduleTestDriver) NewScheduleService() *schedule_application.BookingSrv {
// 	professors := schedule_infra.NewInMemProfRepo(td.professor, false)
// 	students := schedule_infra.NewInMemStudentRepo(td.student, false)
// 	return schedule_application.NewBookingSrv(professors, students)
// }

// func (td *ScheduleTestDriver) BuildStudentContext() context.Context {
// 	ctx := context.Background()
// 	return context.WithValue(ctx, auth_jwt.UserId, studentUuid)
// }
