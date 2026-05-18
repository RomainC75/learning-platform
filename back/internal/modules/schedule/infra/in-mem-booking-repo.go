package schedule_infra

import (
	"errors"
	schedule_domain "language-learning/internal/modules/schedule/domain"
)

type InMemBookingRepo struct {
	bookings []schedule_domain.Booking
}

func NewInMemBookingRepo() *InMemBookingRepo {
	return &InMemBookingRepo{
		bookings: []schedule_domain.Booking{},
	}
}
func (imbr *InMemBookingRepo) isBookable(newBooking schedule_domain.Booking) bool {
	for _, old_booking := range imbr.bookings {
		reza_end := newBooking.GetEndDate()
		old_booking_end := old_booking.GetEndDate()
		if (newBooking.GetStartDate().After(old_booking.GetStartDate()) && newBooking.GetStartDate().Before(old_booking_end)) || (reza_end.After(old_booking.GetStartDate()) && reza_end.Before(old_booking_end)) {
			return false
		}
	}
	return true
}

func (imbr *InMemBookingRepo) SetBooking(newBooking schedule_domain.Booking) error {
	isSchedulable := imbr.isBookable(newBooking)
	if !isSchedulable {
		return errors.New(schedule_domain.ErrNotBookable)
	}
	imbr.bookings = append(imbr.bookings, newBooking)
	return nil
}

func (imbr *InMemBookingRepo) MustAddBooking(newReza schedule_domain.Booking) {
	imbr.bookings = append(imbr.bookings, newReza)
}

func (imbr *InMemBookingRepo) ResetBooking() {
	imbr.bookings = []schedule_domain.Booking{}
}

func (imbr *InMemBookingRepo) ListBookings() []schedule_domain.Booking {
	return imbr.bookings
}
