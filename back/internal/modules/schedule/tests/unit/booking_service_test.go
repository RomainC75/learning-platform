package schedule_unit

import (
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CreateBookingCases struct {
	Info                 string
	CreateBookingRequest dtos.CreateBookingRequest
	ProfessorBookings    []schedule_domain.Booking
	ExpectedResponse     schedule_application.CreateBookingResponse
	IsError              bool
	ErrorMessage         string
}

var bookingTestCases []CreateBookingCases = []CreateBookingCases{
	{
		Info: "student should make Booking",
		CreateBookingRequest: dtos.CreateBookingRequest{
			ProfessorId: uuid.UUID{},
			Date:        studentDate1,
			Duration:    duration1h,
		},
		ProfessorBookings: []schedule_domain.Booking{
			schedule_domain.NewBooking(
				professorDate,
				duration1h,
				otherStudentUuid,
			),
		},
		ExpectedResponse: schedule_application.CreateBookingResponse{
			Status:      true,
			StudentId:   studentUuid,
			ProfessorId: professorUuid,
			Date:        studentDate1,
			Duration:    duration1h,
		},
	},
}

func TestCreateBooking(t *testing.T) {
	for _, cs := range bookingTestCases {
		t.Run(cs.Info, func(t *testing.T) {
			td := NewBookingTestDriver().CreateBookingsForProfessor(cs.ProfessorBookings)
			srv := td.NewScheduleService()

			studentContext := td.BuildStudentContext()
			res, err := srv.CreateBooking(studentContext, cs.CreateBookingRequest)
			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, schedule_domain.ErrProfessorNotSchedulable)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
			}
		})
	}
}
