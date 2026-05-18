package schedule_unit

import (
	"fmt"
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CreateBookingCase struct {
	Info                 string
	CreateBookingRequest dtos.CreateBookingRequest
	Bookings             []schedule_domain.Booking
	ExpectedResponse     schedule_application.CreateBookingResponse
	IsError              bool
	ErrorMessage         string
}

var (
	newBooking = schedule_domain.NewBooking(
		bookingUuid,
		professorDate,
		duration1h,
		professor,
		student,
	)
	newBookings = []schedule_domain.Booking{
		newBooking,
	}

	bookingTestCases []CreateBookingCase = []CreateBookingCase{
		{
			Info: "student should make Booking",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: uuid.UUID{},
				Date:        professorDate,
				Duration:    duration1h,
			},
			Bookings: newBookings,
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        professorDate,
				Duration:    duration1h,
			},
		},
	}
)

func TestCreateBooking(t *testing.T) {
	for _, cs := range bookingTestCases {
		t.Run(cs.Info, func(t *testing.T) {
			td := NewBookingTestDriver()
			fmt.Println("===> BEFORE : ", td.SavedBookings())
			_ = td.NewScheduleService()

			studentContext := td.BuildStudentContext()
			res, err := td.CreateBooking(studentContext, cs.CreateBookingRequest)
			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, schedule_domain.ErrProfessorNotSchedulable)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
				for i := range newBookings {
					assert.Equal(t, newBookings[i].ToSnapshot(), td.SavedBookings()[i].ToSnapshot())
				}
			}
		})
	}
}
