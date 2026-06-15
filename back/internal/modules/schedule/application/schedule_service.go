package schedule_application

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
	dtos_responses "language-learning/internal/api/dtos/responses"
	"language-learning/internal/api/middlewares"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"

	"github.com/google/uuid"
)

type ScheduleSrv struct {
	professors schedule_domain.Professors
}

func NewScheduleSrv(professors schedule_domain.Professors) *ScheduleSrv {
	return &ScheduleSrv{
		professors: professors,
	}
}

func (ScheduleSrv *ScheduleSrv) CreateSchedule(ctx context.Context, createBookingRequest dtos.CreateScheduleRequest) (dtos_responses.CreateScheduleResponse, error) {
	professorId := ctx.Value(middlewares.UserIdKey).(uuid.UUID)
	foundProfessor, err := ScheduleSrv.professors.Get(professorId)
	if err != nil {
		return dtos_responses.CreateScheduleResponse{}, err
	}

	newSchedule := NewScheduleFromDto(createBookingRequest)
	foundProfessor.SetSchedule(newSchedule)

	return dtos_responses.ToCreateScheduleResponse(newSchedule), nil
}

func NewScheduleFromDto(createBookingRequest dtos.CreateScheduleRequest) schedule_domain.Schedule {
	weekAvailabilities := make([]schedule_domain.WeeklyAvailability, 0, len(createBookingRequest.WeeklyAvailabilities))
	for _, wa := range createBookingRequest.WeeklyAvailabilities {
		localTime, _ := shared_domain_time.NewLocalTime24(wa.TimeRange.LocalTimeStart.Hour, wa.TimeRange.LocalTimeStart.Minute)
		timeRange := shared_domain_time.NewTimeRange(localTime, wa.TimeRange.Duration)
		weekAvailability := schedule_domain.NewWeeklyAvailability(wa.Day, timeRange)
		weekAvailabilities = append(weekAvailabilities, weekAvailability)
	}

	exceptions := make([]schedule_domain.AvailabilityException, 0, len(createBookingRequest.AvailabilityExceptions))
	for _, ae := range createBookingRequest.AvailabilityExceptions {
		dateRangeTime := shared_domain_time.NewDateTimeRange(ae.DateTimeRange.Start, ae.DateTimeRange.Duration)
		exception := schedule_domain.NewAvailabilityException(dateRangeTime)
		exceptions = append(exceptions, exception)
	}
	return schedule_domain.NewSchedule(weekAvailabilities, exceptions)
}
