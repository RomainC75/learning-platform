package schedule_domain

import shared_domain_time "language-learning/internal/modules/shared/domain/time"

type AvailabilityException struct {
	dateTimeRange shared_domain_time.DateTimeRange
}

func NewAvailabilityException(dateTimeRange shared_domain_time.DateTimeRange) AvailabilityException {
	return AvailabilityException{
		dateTimeRange: dateTimeRange,
	}
}

func (ae AvailabilityException) IsOverlapping(booking Booking) bool {
	return booking.dateTimeRange.IsOverlapWith(ae.dateTimeRange)
}

// == Snapshot ==

func (ae AvailabilityException) ToSnapshot() AvailabilityExceptionSnapshot {
	return AvailabilityExceptionSnapshot{
		DateTimeRange: ae.dateTimeRange.ToSnapshot(),
	}
}

type AvailabilityExceptionSnapshot struct {
	DateTimeRange shared_domain_time.DateTimeRangeSnapshot
}

func NewAvailabilityExceptionFromSnapshot(snap AvailabilityExceptionSnapshot) AvailabilityException {
	return NewAvailabilityException(shared_domain_time.NewDateTimeRangeFromSnapshot(snap.DateTimeRange))
}
