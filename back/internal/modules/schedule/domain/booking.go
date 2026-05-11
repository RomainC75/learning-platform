package schedule_domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	date     time.Time
	duration time.Duration
	with     uuid.UUID
}

func NewBooking(date time.Time, duration time.Duration, with uuid.UUID) Booking {
	return Booking{
		date:     date,
		duration: duration,
		with:     with,
	}
}

func (booking Booking) isBookableIn(schedule []Booking) bool {
	for _, sch := range schedule {
		reza_end := booking.date.Add(booking.duration)
		sch_end := sch.date.Add(sch.duration)
		if (booking.date.After(sch.date) && booking.date.Before(sch_end)) || (reza_end.After(sch.date) && reza_end.Before(sch_end)) {
			return false
		}
	}
	return true
}
