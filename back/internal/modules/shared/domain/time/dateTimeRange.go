package shared_domain_time

import (
	"time"
)

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

func (dtr DateTimeRange) EndDate() time.Time {
	return dtr.start.Add(dtr.duration)
}

func (dtr DateTimeRange) StartDate() time.Time {
	return dtr.start
}

func (dtr DateTimeRange) StartInMinutesAfter00h() int {
	return dtr.start.Hour()*60 + dtr.start.Minute()
}

func (dtr DateTimeRange) EndInMinutesAfter00h() int {
	return dtr.start.Add(dtr.duration).Hour()*60 + dtr.start.Add(dtr.duration).Minute()
}

func (tdr DateTimeRange) IsOverlapWith(other DateTimeRange) bool {
	otherEnd := other.start.Add(other.duration)

	tdrStartsDuringOther := tdr.startsAfterOrEqual(other.start) && tdr.StartsBeforeOrEqual(otherEnd)
	tdrEndsDuringOther := tdr.endsAfterOrEqual(other.start) && tdr.endsBeforeOrEqual(otherEnd)
	return tdr.starsTogether(other) || tdrStartsDuringOther || tdrEndsDuringOther
}

func (tdr DateTimeRange) starsTogether(otherDateTime DateTimeRange) bool {
	return tdr.start.Equal(otherDateTime.start)
}

func (tdr DateTimeRange) StartsBeforeOrEqual(otherTime time.Time) bool {
	return tdr.start.Compare(otherTime) < 0
}

func (tdr DateTimeRange) startsAfterOrEqual(otherTime time.Time) bool {
	return tdr.start.Compare(otherTime) > 0
}

func (tdr DateTimeRange) endsBeforeOrEqual(otherTime time.Time) bool {
	tdrEnd := tdr.start.Add(tdr.duration)
	return tdrEnd.Compare(otherTime) < 0
}

func (tdr DateTimeRange) endsAfterOrEqual(otherTime time.Time) bool {
	tdrEnd := tdr.start.Add(tdr.duration)
	return tdrEnd.Compare(otherTime) > 0
}

// == Snapshot ==

func (dtr DateTimeRange) ToSnapshot() DateTimeRangeSnapshot {
	return DateTimeRangeSnapshot{
		Start:    dtr.start,
		Duration: dtr.duration,
	}
}

func NewDateTimeRangeFromSnapshot(snap DateTimeRangeSnapshot) DateTimeRange {
	return NewDateTimeRange(snap.Start, snap.Duration)
}
