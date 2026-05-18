package schedule_domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	id        uuid.UUID
	date      time.Time
	duration  time.Duration
	professor *Professor
	student   *Student
}

type BookingSnapshot struct {
	Id        uuid.UUID
	Date      time.Time
	Duration  time.Duration
	Professor *Professor
	Student   *Student
}

func NewBooking(id uuid.UUID, date time.Time, duration time.Duration, professor *Professor, student *Student) Booking {
	return Booking{
		id:        id,
		date:      date,
		duration:  duration,
		professor: professor,
		student:   student,
	}
}

func (b *Booking) GetEndDate() time.Time {
	return b.date.Add(b.duration)
}

func (b *Booking) GetStartDate() time.Time {
	return b.date
}

func (b *Booking) ToSnapshot() *BookingSnapshot {
	return &BookingSnapshot{
		Id:        b.id,
		Date:      b.date,
		Duration:  b.duration,
		Professor: b.professor,
		Student:   b.student,
	}
}
