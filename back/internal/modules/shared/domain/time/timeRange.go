package shared_domain_time

import (
	"time"
)

type TimeRange struct {
	localTimeStart LocalTime
	duration       time.Duration
}

type TimeRangeSnapshot struct {
	LocalTimeStart LocalTimeSnapshot
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
		LocalTimeStart: tr.localTimeStart.ToSnapshot(),
		Duration:       tr.duration,
	}
}

func (tr TimeRange) IsContaining(other DateTimeRange) bool {
	trStartMinAfter00h := tr.localTimeStart.StartInMinutesAfter00h()
	trEndMinAfter00h := trStartMinAfter00h + int(tr.duration.Minutes())
	return trStartMinAfter00h <= other.StartInMinutesAfter00h() && trStartMinAfter00h <= other.EndInMinutesAfter00h() && trEndMinAfter00h >= other.StartInMinutesAfter00h() && trEndMinAfter00h >= other.EndInMinutesAfter00h()
}

func NewTimeRangeFromSnapshot(snap TimeRangeSnapshot) TimeRange {
	return NewTimeRange(NewLocalTimeFromSnapshot(snap.LocalTimeStart), snap.Duration)
}
