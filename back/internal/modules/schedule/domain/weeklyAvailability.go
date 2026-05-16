package schedule_domain

import (
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"time"
)

type WeeklyAvailability struct {
	day       time.Weekday
	timeRange shared_domain_time.TimeRange
}

type WeeklyAvailabilitySnapshot struct {
	Day       time.Weekday
	TimeRange shared_domain_time.TimeRangeSnapshot
}

func NewWeeklyAvailability(day time.Weekday, timeRange shared_domain_time.TimeRange) WeeklyAvailability {
	return WeeklyAvailability{
		day:       day,
		timeRange: timeRange,
	}
}

func (wa WeeklyAvailability) ToSnapshot() WeeklyAvailabilitySnapshot {
	return WeeklyAvailabilitySnapshot{
		Day:       wa.day,
		TimeRange: wa.timeRange.ToSnapshot(),
	}
}
