package shared_domain_time

import "time"

type DateTimeRange struct {
	start    time.Time
	duration time.Duration
}

type DateTimeRangeSnapshot struct {
	Start    time.Time
	Duration time.Duration
}

func NewDateTimeRange(start time.Time, duration time.Duration) DateTimeRange {
	return DateTimeRange{
		start:    start,
		duration: duration,
	}
}

func (dtr DateTimeRange) ToSnapshot() DateTimeRangeSnapshot {
	return DateTimeRangeSnapshot{
		Start:    dtr.start,
		Duration: dtr.duration,
	}
}
