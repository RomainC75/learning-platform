package schedule_domain

import (
	"errors"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBookingInPast = "error : cannot create booking in the past"
)

type Booking struct {
	id            uuid.UUID
	dateTimeRange shared_domain_time.DateTimeRange
	professor     ProfessorId
	student       StudentId
	createdAt     time.Time
}

func NewBooking(id uuid.UUID, dateTimeRange shared_domain_time.DateTimeRange, professor ProfessorId, student StudentId, createdAt time.Time) (Booking, error) {
	if dateTimeRange.StartsBeforeOrEqual(createdAt) {
		return Booking{}, errors.New(ErrBookingInPast)
	}
	return Booking{
		id:            id,
		dateTimeRange: dateTimeRange,
		professor:     professor,
		student:       student,
		createdAt:     createdAt,
	}, nil
}

func MustNewBooking(id uuid.UUID, dateTimeRange shared_domain_time.DateTimeRange, professor ProfessorId, student StudentId, createdAt time.Time) Booking {
	newBooking, err := NewBooking(id, dateTimeRange, professor, student, createdAt)
	if err != nil {
		panic(err)
	}
	return newBooking
}

func (b *Booking) EndDate() time.Time {
	return b.dateTimeRange.EndDate()
}

func (b *Booking) StartDate() time.Time {
	return b.dateTimeRange.StartDate()
}

func (b *Booking) DateTimeRange() shared_domain_time.DateTimeRange {
	return b.dateTimeRange
}

func isBookingAlreadyExists(bookingList []Booking, newBooking Booking) bool {
	for _, booking := range bookingList {
		if booking.dateTimeRange.IsOverlapWith(newBooking.dateTimeRange) {
			return true
		}
	}
	return false
}

// == Snapshot ==

type BookingSnapshot struct {
	Id            uuid.UUID
	DateTimeRange shared_domain_time.DateTimeRangeSnapshot
	Professor     ProfessorId
	Student       StudentId
	CreatedAt     time.Time
}

func (b *Booking) ToSnapshot() *BookingSnapshot {
	return &BookingSnapshot{
		Id:            b.id,
		DateTimeRange: b.dateTimeRange.ToSnapshot(),
		Professor:     b.professor,
		Student:       b.student,
		CreatedAt:     b.createdAt,
	}
}
