package schedule_unit

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
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
	ExpectedResponse         string
}

var testCases []CreateReservationCases = []CreateReservationCases{
	{
		CreateReservationRequest: dtos.CreateReservationRequest{
			ProfessorId: uuid.UUID{},
			Date:        time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC),
			Duration:    utils_time.MustParseDuration("1h"),
		},
		ProfessorReservations: []schedule_domain.Reservation{
			schedule_domain.NewReservation(
				time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC),
				utils_time.MustParseDuration("1h"),
			),
		},
		ExpectedResponse: "scheduled",
	},
}

func TestCreateReservation(t *testing.T) {
	for _, cs := range testCases {
		srv := NewTestDriver().CreateReservationsForProfessor(cs.ProfessorReservations).NewScheduleService()

		ctx := context.Background()
		res, err := srv.CreateReservation(ctx, cs.CreateReservationRequest)
		assert.Nil(t, err)
		assert.Equal(t, cs.ExpectedResponse, res)
	}
}
