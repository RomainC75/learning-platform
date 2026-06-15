package schedule_domain

import (
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"time"
)

type WeeklyAvailability struct {
	day       time.Weekday
	timeRange shared_domain_time.TimeRange
}

func NewWeeklyAvailability(day time.Weekday, timeRange shared_domain_time.TimeRange) WeeklyAvailability {
	return WeeklyAvailability{
		day:       day,
		timeRange: timeRange,
	}
}

func (wa WeeklyAvailability) IsMatchWith(newBooking Booking) bool {
	return wa.day == newBooking.StartDate().Weekday() && wa.timeRange.IsContaining(newBooking.DateTimeRange())
}

// == Snapshot ==

type WeeklyAvailabilitySnapshot struct {
	Day       time.Weekday
	TimeRange shared_domain_time.TimeRangeSnapshot
}

func (wa WeeklyAvailability) ToSnapshot() WeeklyAvailabilitySnapshot {
	return WeeklyAvailabilitySnapshot{
		Day:       wa.day,
		TimeRange: wa.timeRange.ToSnapshot(),
	}
}

func NewWeeklyAvailabilityFromSnapshot(snap WeeklyAvailabilitySnapshot) WeeklyAvailability {
	return NewWeeklyAvailability(snap.Day, shared_domain_time.NewTimeRangeFromSnapshot(snap.TimeRange))
}
