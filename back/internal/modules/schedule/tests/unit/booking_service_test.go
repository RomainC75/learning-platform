package schedule_unit

import (
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
	controlBooking = schedule_domain.NewBooking(
		bookingUuid,
		professorDate,
		duration1h,
		professor,
		student,
		nowjuly1,
	)
	controlBookings = []schedule_domain.Booking{
		controlBooking,
	}

	bookingTestCases []CreateBookingCase = []CreateBookingCase{
		{
			Info: "student should make Booking",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: uuid.UUID{},
				Date:        professorDate,
				Duration:    duration1h,
			},
			Bookings: controlBookings,
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
			_ = td.NewScheduleService()

			studentContext := td.BuildStudentContext()
			res, err := td.CreateBooking(studentContext, cs.CreateBookingRequest)
			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, schedule_domain.ErrProfessorNotSchedulable)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
				for i := range controlBookings {
					assert.Equal(t, controlBookings[i].ToSnapshot(), td.SavedBookings()[i].ToSnapshot())
				}
			}
		})
	}
}
