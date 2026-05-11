package schedule_unit

import (
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CreateReservationCases struct {
	Info                     string
	CreateReservationRequest dtos.CreateReservationRequest
	ProfessorReservations    []schedule_domain.Reservation
	ExpectedResponse         schedule_application.CreateReservationResponse
	IsError                  bool
	ErrorMessage             string
}

var (
	studentDate1  = time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC)
	studentDate2  = time.Date(2026, time.May, 9, 10, 30, 0, 0, time.UTC)
	professorDate = time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC)
	duration1h    = utils_time.MustParseDuration("1h")
)

var testCases []CreateReservationCases = []CreateReservationCases{
	{
		Info: "should make reservation",
		CreateReservationRequest: dtos.CreateReservationRequest{
			ProfessorId: uuid.UUID{},
			Date:        studentDate1,
			Duration:    duration1h,
		},
		ProfessorReservations: []schedule_domain.Reservation{
			schedule_domain.NewReservation(
				professorDate,
				duration1h,
				otherStudentUuid,
			),
		},
		ExpectedResponse: schedule_application.CreateReservationResponse{
			Status:      true,
			StudentId:   studentUuid,
			ProfessorId: professorUuid,
			Date:        studentDate1,
			Duration:    duration1h,
		},
	},
	{
		Info: "should NOT make reservation",
		CreateReservationRequest: dtos.CreateReservationRequest{
			ProfessorId: uuid.UUID{},
			Date:        studentDate2,
			Duration:    duration1h,
		},
		ProfessorReservations: []schedule_domain.Reservation{
			schedule_domain.NewReservation(
				professorDate,
				duration1h,
				otherStudentUuid,
			),
		},
		IsError:      true,
		ErrorMessage: schedule_domain.ErrProfessorNotSchedulable,
	},
}

func TestCreateReservation(t *testing.T) {
	for _, cs := range testCases {
		t.Run(cs.Info, func(t *testing.T) {
			td := NewTestDriver().CreateReservationsForProfessor(cs.ProfessorReservations)
			srv := td.NewScheduleService()

			studentContext := td.BuildStudentContext()
			res, err := srv.CreateReservation(studentContext, cs.CreateReservationRequest)
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
