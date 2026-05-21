package schedule_unit

import (
	"fmt"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type AETestCase struct {
	Message  string
	Date     time.Time
	Duration time.Duration
	Booking  schedule_domain.Booking
	Expected bool
}

var (
	aeControlBooking = schedule_domain.MustNewBooking(
		booked1Uuid,
		shared_domain_time.NewDateTimeRange(date_2026_july_6_12h00, duration1h),
		professor,
		student,
		nowMarch1,
	)
	aeTestCases = []AETestCase{
		{
			Message:  "should overlap availability exception because inside",
			Date:     july1,
			Duration: duration1Week,
			Booking:  aeControlBooking,
			Expected: true,
		},
		{
			Message: "should overlap availability exception because booking starts during exception",
			Date:    july1,
			// end : july 8 : 12h30
			Duration: duration1Week + 12*time.Hour + 30*time.Minute,
			Booking:  aeControlBooking,
			Expected: true,
		},
		{
			Message:  "should overlap availability exception because booking ends during exception",
			Date:     july1,
			Duration: 5*time.Hour*24 + time.Hour*12 + 30*time.Minute,
			Booking:  aeControlBooking,
			Expected: true,
		},
		{
			Message:  "should NOT overlap availability exception because booking is BEFORE exception",
			Date:     july1,
			Duration: duration4h,
			Booking:  aeControlBooking,
			Expected: false,
		},
		{
			Message:  "should NOT overlap availability exception because booking is AFTER exception",
			Date:     july25,
			Duration: duration1Week,
			Booking:  aeControlBooking,
			Expected: false,
		},
	}
)

func TestAvailabilityException(t *testing.T) {
	for _, tc := range aeTestCases {
		t.Run(tc.Message, func(t *testing.T) {
			fmt.Println("====> ends : ", july1.Add(tc.Duration).UTC())
			fmt.Println("  ")

			dtr := shared_domain_time.NewDateTimeRange(tc.Date, tc.Duration)

			ae := schedule_domain.NewAvailabilityException(dtr)

			assert.Equal(t, tc.Expected, ae.IsOverlapping(tc.Booking))
		})
	}
}
