package dtos_responses

import (
	"language-learning/internal/api/dtos"
	schedule_domain "language-learning/internal/modules/schedule/domain"
)

type CreateScheduleResponse struct {
	WeeklyAvailabilities   []dtos.WeeklyAvailabilityDTO    `json:"weekly_availabilities"`
	AvailabilityExceptions []dtos.AvailabilityExceptionDTO `json:"exceptions"`
}

func ToCreateScheduleResponse(
	schedule schedule_domain.Schedule,
) CreateScheduleResponse {
	scheduleSnapshot := schedule.ToSnapshot()

	weeklyAvailabilities := make([]dtos.WeeklyAvailabilityDTO, len(scheduleSnapshot.WeeklyAvailabilities))
	for i, wa := range scheduleSnapshot.WeeklyAvailabilities {
		weeklyAvailabilities[i] = dtos.WeeklyAvailabilityDTO{
			Day: wa.Day,
			TimeRange: dtos.TimeRangeDTO{
				LocalTimeStart: dtos.LocalTimeDTO{
					Hour:   wa.TimeRange.LocalTimeStart.ToSnapshot().Hour,
					Minute: wa.TimeRange.LocalTimeStart.ToSnapshot().Minute,
				},
				Duration: wa.TimeRange.Duration,
			},
		}
	}

	exceptions := make([]dtos.AvailabilityExceptionDTO, len(scheduleSnapshot.AvailabilityExceptions))
	for i, ae := range scheduleSnapshot.AvailabilityExceptions {
		exceptions[i] = dtos.AvailabilityExceptionDTO{
			DateTimeRange: dtos.DateTimeRangeDTO{
				Start:    ae.DateTimeRange.Start,
				Duration: ae.DateTimeRange.Duration,
			},
		}
	}

	return CreateScheduleResponse{
		WeeklyAvailabilities:   weeklyAvailabilities,
		AvailabilityExceptions: exceptions,
	}
}
