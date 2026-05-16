package dtos_requests

import "language-learning/internal/api/dtos"

type CreateScheduleRequest struct {
	WeeklyAvailabilities   []dtos.WeeklyAvailabilityDTO    `json:"weekly_availabilities"`
	AvailabilityExceptions []dtos.AvailabilityExceptionDTO `json:"exceptions"`
}
