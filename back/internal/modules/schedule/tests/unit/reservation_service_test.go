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
	CreateReservationRequest dtos.CreateReservationRequest
	ProfessorReservations    []schedule_domain.Reservation
	ExpectedResponse         schedule_application.CreateReservationResponse
}

var (
	studentDate   = time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC)
	professorDate = time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC)
	duration1h    = utils_time.MustParseDuration("1h")
)

var testCases []CreateReservationCases = []CreateReservationCases{
	{
		CreateReservationRequest: dtos.CreateReservationRequest{
			ProfessorId: uuid.UUID{},
			Date:        studentDate,
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
			Date:        studentDate,
			Duration:    duration1h,
		},
	},
}

func TestCreateReservation(t *testing.T) {
	for _, cs := range testCases {
		td := NewTestDriver().CreateReservationsForProfessor(cs.ProfessorReservations)
		srv := td.NewScheduleService()

		studentContext := td.BuildStudentContext()
		res, err := srv.CreateReservation(studentContext, cs.CreateReservationRequest)
		assert.Nil(t, err)
		assert.Equal(t, cs.ExpectedResponse, res)
	}
}
