package schedule_domain

import (
	"slices"
)

type Schedule struct {
	weeklyAvailabilities   []WeeklyAvailability
	availabilityExceptions []AvailabilityException
}

func NewSchedule(weeklyAvailabilities []WeeklyAvailability, exceptions []AvailabilityException) Schedule {
	return Schedule{
		weeklyAvailabilities:   weeklyAvailabilities,
		availabilityExceptions: exceptions,
	}
}

func (s Schedule) WeeklyAvailabilities() []WeeklyAvailability {
	return slices.Clone(s.weeklyAvailabilities)
}

func (s Schedule) AvailabilityExceptions() []AvailabilityException {
	return slices.Clone(s.availabilityExceptions)
}

func (s Schedule) IsAMatchWithWeeklyAvailabilities(newBooking Booking) bool {
	for _, wa := range s.weeklyAvailabilities {
		if wa.IsMatchWith(newBooking) {
			return true
		}
	}
	return false
}

func (s Schedule) IsNotRespectingAvailabilityExceptions(booking Booking) bool {
	for _, availabilityException := range s.availabilityExceptions {
		if availabilityException.IsOverlapping(booking) {
			return true
		}
	}
	return false
}

// == Snapshot ==

type ScheduleSnapshot struct {
	WeeklyAvailabilities   []WeeklyAvailabilitySnapshot
	AvailabilityExceptions []AvailabilityExceptionSnapshot
}

func (s Schedule) ToSnapshot() ScheduleSnapshot {
	weeklyAvailabilities := make([]WeeklyAvailabilitySnapshot, len(s.weeklyAvailabilities))
	for i, wa := range s.weeklyAvailabilities {
		weeklyAvailabilities[i] = wa.ToSnapshot()
	}

	availabilityExceptions := make([]AvailabilityExceptionSnapshot, len(s.availabilityExceptions))
	for i, ae := range s.availabilityExceptions {
		availabilityExceptions[i] = ae.ToSnapshot()
	}

	return ScheduleSnapshot{
		WeeklyAvailabilities:   weeklyAvailabilities,
		AvailabilityExceptions: availabilityExceptions,
	}
}

func NewScheduleFromSnapshot(snap ScheduleSnapshot) Schedule {
	weeklyAvailabilities := make([]WeeklyAvailability, 0, len(snap.WeeklyAvailabilities))
	for _, waSnap := range snap.WeeklyAvailabilities {
		wa := NewWeeklyAvailabilityFromSnapshot(waSnap)
		weeklyAvailabilities = append(weeklyAvailabilities, wa)
	}

	availabilityExceptions := make([]AvailabilityException, 0, len(snap.AvailabilityExceptions))
	for _, aeSnap := range snap.AvailabilityExceptions {
		ae := NewAvailabilityExceptionFromSnapshot(aeSnap)
		availabilityExceptions = append(availabilityExceptions, ae)
	}

	return NewSchedule(weeklyAvailabilities, availabilityExceptions)
}
