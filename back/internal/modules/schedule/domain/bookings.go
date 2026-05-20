package schedule_domain

var (
	ErrNotBookable = "not bookable"
)

type Bookings interface {
	SetBooking(booking Booking) error
	GetBookings(professor *Professor) []Booking
	ResetBooking()
}
