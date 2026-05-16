package dtos

import "time"

type WeeklyAvailabilityDTO struct {
	Day       time.Weekday `json:"day"`
	TimeRange TimeRangeDTO `json:"time_range"`
}
type TimeRangeDTO struct {
	LocalTimeStart LocalTimeDTO  `json:"local_time_start"`
	Duration       time.Duration `json:"duration"`
}

type DateRangeTimeRangeDTO struct {
	Start    time.Time     `json:"start"`
	Duration time.Duration `json:"duration"`
}

type LocalTimeDTO struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type AvailabilityExceptionDTO struct {
	DateTimeRange DateTimeRangeDTO `json:"date_time_range"`
}

type DateTimeRangeDTO struct {
	Start    time.Time     `json:"start"`
	Duration time.Duration `json:"duration"`
}
