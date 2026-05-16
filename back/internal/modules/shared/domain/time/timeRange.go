package shared_domain_time

import "time"

type TimeRange struct {
	localTimeStart LocalTime
	duration       time.Duration
}

type TimeRangeSnapshot struct {
	LocalTimeStart LocalTime
	Duration       time.Duration
}

func NewTimeRange(localTimeStart LocalTime, duration time.Duration) TimeRange {
	return TimeRange{
		localTimeStart: localTimeStart,
		duration:       duration,
	}
}

func (tr TimeRange) ToSnapshot() TimeRangeSnapshot {
	return TimeRangeSnapshot{
		LocalTimeStart: tr.localTimeStart,
		Duration:       tr.duration,
	}
}
