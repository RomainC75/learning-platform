package schedule_unit

import (
	"fmt"
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CreateBookingCase struct {
	Info                 string
	CreateBookingRequest dtos.CreateBookingRequest
	ExpectedResponse     schedule_application.CreateBookingResponse
	IsError              bool
	ErrorMessage         string
}

var (
	controlBooking = schedule_domain.NewBooking(
		booked1Uuid,
		shared_domain_time.NewDateTimeRange(date_2026_may_20_12h00, duration1h),
		professor,
		student,
		nowMarch1,
	)
	bookingTestCases []CreateBookingCase = []CreateBookingCase{
		{
			Info: "should not create booking, cause weeklyavailability",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: professorUuid,
				Date:        date_2026_may_20_12h00,
				Duration:    duration1h,
			},
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        date_2026_may_20_12h00,
				Duration:    duration1h,
			},
			IsError:      true,
			ErrorMessage: schedule_domain.ErrNotMatchWithWeeklyAvailabilities,
		},
		{
			Info: "should not create booking, cause availabilityException",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: professorUuid,
				Date:        date_2026_july_6_12h00,
				Duration:    duration1h,
			},
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        date_2026_july_6_12h00,
				Duration:    duration1h,
			},
			IsError:      true,
			ErrorMessage: schedule_domain.ErrNotRespectingAvailabilityExceptions,
		},
	}
)

func TestCreateBooking(t *testing.T) {
	for _, cs := range bookingTestCases {
		t.Run(cs.Info, func(t *testing.T) {
			td := NewBookingTestDriver()
			td.CreateBookingsForProfessor(bookedList)
			_ = td.NewScheduleService()

			studentContext := td.BuildStudentContext()
			res, err := td.CreateBooking(studentContext, cs.CreateBookingRequest)
			fmt.Println("====> res : ", res)
			if cs.IsError {
				fmt.Println("======= ERR : ", err.Error())
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, cs.ErrorMessage)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)

				controlBookings := make([]schedule_domain.Booking, len(bookedList)+1)
				controlBookings = append(controlBookings, bookedList...)
				controlBookings = append(controlBookings, controlBooking)

				for i := range controlBookings {
					assert.Equal(t, controlBookings[i].ToSnapshot(), td.SavedBookings()[i].ToSnapshot())
				}
			}
		})
	}
}
