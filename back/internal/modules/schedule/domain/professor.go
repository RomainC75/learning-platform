package schedule_domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrProfessorNotSchedulable             = "professor not schedulable"
	ErrNotMatchWithWeeklyAvailabilities    = "no match with weekly availabilities"
	ErrNotRespectingAvailabilityExceptions = "not respecting availability exceptions"
	ErrBookingAlreadyExists                = "booking already exists"
)

type Professor struct {
	id        uuid.UUID
	firstName string
	lastName  string
	schedule  Schedule
}

func NewProfessor(id uuid.UUID, firstName string, lastName string, schedule Schedule) *Professor {
	return &Professor{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		schedule:  schedule,
	}
}

func (p *Professor) SetSchedule(schedule Schedule) {
	p.schedule = schedule
}

func (p *Professor) Id() uuid.UUID {
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
	fmt.Println("===============> NNNOOO ERROR ")
	for _, b := range bookingList {
		fmt.Println("===============> BOOKING IN LIST", b.ToSnapshot())
	}
	fmt.Println("======================================> NEW BOOKIN", newBooking.ToSnapshot())
	return nil

}
