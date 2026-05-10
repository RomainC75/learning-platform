package unit_test

import (
	"context"
	dtos "language-learning/internal/api/dtos/request"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type CreateReservationCases struct {
	CreateReservationRequest dtos.CreateReservationRequest
	Student                  *schedule_domain.Student
	Professor                *schedule_domain.Professor
	ExpectedResponse         string
}

var testCases []CreateReservationCases = []CreateReservationCases{
	{
		CreateReservationRequest: dtos.CreateReservationRequest{
			ProfessorId: uuid.UUID{},
			Date:        time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC),
			Duration:    utils_time.MustParseDuration("1h"),
		},
		Student: schedule_domain.NewStudent(
			uuid.New(),
			"bob",
			"sponge",
			[]schedule_domain.Reservation{},
		),
		Professor: schedule_domain.NewProfessor(
			uuid.UUID{},
			"big",
			"brother",
			[]schedule_domain.Reservation{
				schedule_domain.NewReservation(
					time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC),
					utils_time.MustParseDuration("1h"),
				),
			},
		),
		ExpectedResponse: "scheduled",
	},
}

func TestCreateReservation(t *testing.T) {
	for _, cs := range testCases {
		professors := schedule_infra.NewInMemProfRepo(cs.Professor, false)
		students := schedule_infra.NewInMemStudentRepo(cs.Student, false)
		srv := schedule_application.NewReservationSrv(professors, students)

		ctx := context.Background()
		res, err := srv.CreateReservation(ctx, cs.CreateReservationRequest)
		assert.Nil(t, err)
		assert.Equal(t, cs.ExpectedResponse, res)
	}
}
