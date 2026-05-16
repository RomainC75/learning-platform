package schedule_domain

import shared_domain_time "language-learning/internal/modules/shared/domain/time"

type AvailabilityException struct {
	dateTimeRange shared_domain_time.DateTimeRange
}

type AvailabilityExceptionSnapshot struct {
	DateTimeRange shared_domain_time.DateTimeRangeSnapshot
}

func NewAvailabilityException(dateTimeRange shared_domain_time.DateTimeRange) AvailabilityException {
	return AvailabilityException{
		dateTimeRange: dateTimeRange,
	}
}

func (ae AvailabilityException) ToSnapshot() AvailabilityExceptionSnapshot {
	return AvailabilityExceptionSnapshot{
		DateTimeRange: ae.dateTimeRange.ToSnapshot(),
	}
}
