package schedule_unit

import (
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	utils_display "language-learning/utils/display"
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
	controlBooking = schedule_domain.MustNewBooking(
		newBookingUuid,
		shared_domain_time.NewDateTimeRange(date_2026_may_25_12h00, duration1h),
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
		{
			Info: "should not create booking, cause booking already exists",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: professorUuid,
				Date:        date_2026_may_12_14h00,
				Duration:    duration1h,
			},
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        date_2026_may_12_14h00,
				Duration:    duration1h,
			},
			IsError:      true,
			ErrorMessage: schedule_domain.ErrBookingAlreadyExists,
		},
		{
			Info: "should create booking and save it",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: professorUuid,
				Date:        date_2026_may_25_12h00,
				Duration:    duration1h,
			},
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        date_2026_may_25_12h00,
				Duration:    duration1h,
			},
			IsError: false,
		},
		{
			Info: "should not create booking in the past",
			CreateBookingRequest: dtos.CreateBookingRequest{
				ProfessorId: professorUuid,
				Date:        date_2026_feb_1_12h00,
				Duration:    duration1h,
			},
			ExpectedResponse: schedule_application.CreateBookingResponse{
				Status:      true,
				StudentId:   studentUuid,
				ProfessorId: professorUuid,
				Date:        date_2026_feb_1_12h00,
				Duration:    duration1h,
			},
			IsError:      true,
			ErrorMessage: schedule_domain.ErrBookingInPast,
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
			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, cs.ErrorMessage)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)

				controlBookings := make([]schedule_domain.Booking, 0, len(bookedList)+1)
				controlBookings = append(controlBookings, bookedList...)
				controlBookings = append(controlBookings, controlBooking)

				assert.Equal(t, len(controlBookings), len(td.SavedBookings()))
				for i := range controlBookings {
					utils_display.PrettyDisplay("control : ", controlBookings[i].ToSnapshot())
					utils_display.PrettyDisplay("savedBooking : ", td.SavedBookings()[i].ToSnapshot())
					assert.Equal(t, controlBookings[i].ToSnapshot(), td.SavedBookings()[i].ToSnapshot())
				}

			}
		})
	}
}
