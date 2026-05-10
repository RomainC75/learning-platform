package schedule_domain

import (
	"time"
)

type Reservation struct {
	date     time.Time
	duration time.Duration
}

func NewReservation(date time.Time, duration time.Duration) Reservation {
	return Reservation{
		date:     date,
		duration: duration,
	}
}

func (reservation Reservation) IsReservableIn(schedule []Reservation) bool {
	for _, sch := range schedule {
		reza_end := reservation.date.Add(reservation.duration)
		sch_end := sch.date.Add(sch.duration)
		if (reservation.date.After(sch.date) && reservation.date.Before(sch_end)) || (reza_end.After(sch.date) && reza_end.Before(sch_end)) {
			return false
		}
	}
	return true
}
