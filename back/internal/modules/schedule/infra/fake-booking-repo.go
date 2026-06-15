package schedule_infra

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
)

type FakeBookingRepo struct {
	bookings []schedule_domain.Booking
}

func NewFakeBookingRepo() *FakeBookingRepo {
	return &FakeBookingRepo{
		bookings: []schedule_domain.Booking{},
	}
}

func (imbr *FakeBookingRepo) SetBooking(newBooking schedule_domain.Booking) error {
	// isSchedulable := imbr.isBookable(newBooking)
	// if !isSchedulable {
	// 	return errors.New(schedule_domain.ErrNotBookable)
	// }
	imbr.bookings = append(imbr.bookings, newBooking)
	return nil
}

func (imbr *FakeBookingRepo) GetBookingsByProfessorId(professorId schedule_domain.ProfessorId) []schedule_domain.Booking {
	return imbr.bookings
}

func (imbr *FakeBookingRepo) MustAddBooking(newReza schedule_domain.Booking) {
	imbr.bookings = append(imbr.bookings, newReza)
}

func (imbr *FakeBookingRepo) ResetBooking() {
	imbr.bookings = []schedule_domain.Booking{}
}

func (imbr *FakeBookingRepo) ListBookings() []schedule_domain.Booking {
	return imbr.bookings
}
