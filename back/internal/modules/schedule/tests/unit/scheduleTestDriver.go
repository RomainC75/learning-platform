package schedule_unit

import (
	"context"
	dtos_requests "language-learning/internal/api/dtos/request"
	dtos_responses "language-learning/internal/api/dtos/responses"
	"language-learning/internal/api/middlewares"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	utils_time "language-learning/utils/time"
	"time"
)

type ScheduleTestDriver struct {
	professor   *schedule_domain.Professor
	scheduleSrv *schedule_application.ScheduleSrv
	professors  *schedule_infra.FakeProfRepo
	ctx         context.Context
}

func NewScheduleTestDriver() *ScheduleTestDriver {
	professor := schedule_domain.NewProfessor(professorUuid, schedule_domain.NewSchedule(
		[]schedule_domain.WeeklyAvailability{
			schedule_domain.NewWeeklyAvailability(
				time.Monday,
				shared_domain_time.NewTimeRange(
					shared_domain_time.MustLocalTime24(9, 0),
					utils_time.MustParseDuration("4h"),
				),
			),
		},
		[]schedule_domain.AvailabilityException{
			schedule_domain.NewAvailabilityException(
				shared_domain_time.NewDateTimeRange(
					july25,
					duration1Week,
				),
			),
		},
	))

	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.UserIdKey, professorUuid)
	return &ScheduleTestDriver{
		professor: professor,
		ctx:       ctx,
	}

}

func (td *ScheduleTestDriver) NewScheduleService(isProfessorMissing bool) {
	professors := schedule_infra.NewFakeProfRepo(td.professor, isProfessorMissing)
	td.professors = professors
	td.scheduleSrv = schedule_application.NewScheduleSrv(professors)
}

func (td *ScheduleTestDriver) SetSchedule(schedule schedule_domain.Schedule) {
	td.professor.SetSchedule(schedule)
}

func (td *ScheduleTestDriver) GetSavedSchedule() schedule_domain.Schedule {
	return td.professors.GetSavedSchedule()
}

func (td *ScheduleTestDriver) CreateSchedule(createSchedule dtos_requests.CreateScheduleRequest) (dtos_responses.CreateScheduleResponse, error) {
	return td.scheduleSrv.CreateSchedule(td.ctx, createSchedule)
}

func (td *ScheduleTestDriver) ResetSchedule() {
	td.professor.SetSchedule(schedule_domain.Schedule{})
}
