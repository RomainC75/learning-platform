package schedule_domain

import "slices"

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

type ScheduleSnapshot struct {
	WeeklyAvailabilities   []WeeklyAvailabilitySnapshot
	AvailabilityExceptions []AvailabilityExceptionSnapshot
}
