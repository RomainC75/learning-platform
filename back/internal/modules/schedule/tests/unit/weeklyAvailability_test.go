package schedule_unit

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type WETestCase struct {
	Message               string
	LocalTimeStartHours   int
	LocalTimeStartMinutes int
	Duration              string
	Day                   time.Weekday
	Booking               schedule_domain.Booking
	Expected              bool
}

var waControlBooking = schedule_domain.MustNewBooking(
	booked1Uuid,
	shared_domain_time.NewDateTimeRange(date_2026_may_11_12h00, duration1h),
	professor.Id(),
	student,
	nowMarch1,
)

var WETestCases []WETestCase = []WETestCase{
	{
		Message:               "should match weekly availability",
		LocalTimeStartHours:   11,
		LocalTimeStartMinutes: 0,
		Duration:              "3h",
		Day:                   time.Monday,
		Booking:               waControlBooking,
		Expected:              true,
	},
	{
		Message:               "should match weekly availability",
		LocalTimeStartHours:   12,
		LocalTimeStartMinutes: 0,
		Duration:              "3h",
		Day:                   time.Monday,
		Booking:               waControlBooking,
		Expected:              true,
	},
	{
		Message:               "should match weekly availability",
		LocalTimeStartHours:   11,
		LocalTimeStartMinutes: 0,
		Duration:              "2h",
		Day:                   time.Monday,
		Booking:               waControlBooking,
		Expected:              true,
	},
	{
		Message:               "should NOT match weekly availability",
		LocalTimeStartHours:   12,
		LocalTimeStartMinutes: 30,
		Duration:              "2h",
		Day:                   time.Monday,
		Booking:               waControlBooking,
		Expected:              false,
	},
	{
		Message:               "should NOT match weekly availability",
		LocalTimeStartHours:   11,
		LocalTimeStartMinutes: 30,
		Duration:              "1h",
		Day:                   time.Monday,
		Booking:               waControlBooking,
		Expected:              false,
	},
}

func TestWeeklyAvailability(t *testing.T) {
	for _, tc := range WETestCases {
		t.Run(tc.Message, func(t *testing.T) {
			tr := shared_domain_time.NewTimeRange(shared_domain_time.MustLocalTime24(tc.LocalTimeStartHours, tc.LocalTimeStartMinutes), utils_time.MustParseDuration(tc.Duration))
			we := schedule_domain.NewWeeklyAvailability(tc.Day, tr)

			assert.Equal(t, we.IsMatchWith(tc.Booking), tc.Expected)
		})
	}
}
