package schedule_domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProfessorNotSchedulable             = "professor not schedulable"
	ErrNotMatchWithWeeklyAvailabilities    = "no match with weekly availabilities"
	ErrNotRespectingAvailabilityExceptions = "not respecting availability exceptions"
	ErrBookingAlreadyExists                = "booking already exists"
)

type ProfessorId uuid.UUID

type Professor struct {
	id       ProfessorId
	schedule Schedule
}

func NewProfessor(userId uuid.UUID, schedule Schedule) *Professor {
	return &Professor{
		id:       ProfessorId(userId),
		schedule: schedule,
	}
}

func (p *Professor) SetSchedule(schedule Schedule) {
	p.schedule = schedule
}

func (p *Professor) Id() ProfessorId {
	return p.id

}

func (p *Professor) IsBookable(bookingList []Booking, newBooking Booking) error {
	if !p.schedule.IsAMatchWithWeeklyAvailabilities(newBooking) {
		return errors.New(ErrNotMatchWithWeeklyAvailabilities)
	} else if p.schedule.IsNotRespectingAvailabilityExceptions(newBooking) {
		return errors.New(ErrNotRespectingAvailabilityExceptions)
	} else if isBookingAlreadyExists(bookingList, newBooking) {
		return errors.New(ErrBookingAlreadyExists)
	}
	return nil
}

// == Snapshot ==

type ProfessorSnapshot struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Schedule  ScheduleSnapshot
}

func (p Professor) ToSnapshot() ProfessorSnapshot {
	return ProfessorSnapshot{
		Id:       uuid.UUID(p.id),
		Schedule: p.schedule.ToSnapshot(),
	}
}
