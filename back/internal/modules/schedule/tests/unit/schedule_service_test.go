package schedule_unit

import (
	"context"
	"fmt"
	"language-learning/internal/api/dtos"
	dtos_requests "language-learning/internal/api/dtos/request"
	dtos_responses "language-learning/internal/api/dtos/responses"
	auth_jwt "language-learning/internal/auth/jwt"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	schedule_infra "language-learning/internal/modules/schedule/infra"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CreateAvailabilityScheduleCases struct {
	Info                  string
	Professor             *schedule_domain.Professor
	CreateScheduleRequest dtos_requests.CreateScheduleRequest
	ProfessorBookings     []schedule_domain.Booking
	ExpectedResponse      dtos_responses.CreateScheduleResponse
	IsError               bool
	ErrorMessage          string
}

var (
	baseSchedule = dtos_requests.CreateScheduleRequest{
		WeeklyAvailabilities: []dtos.WeeklyAvailabilityDTO{
			{
				Day: time.Monday,
				TimeRange: dtos.TimeRangeDTO{
					LocalTimeStart: dtos.LocalTimeDTO{
						Hour:   9,
						Minute: 0,
					},
					Duration: utils_time.MustParseDuration("4h"),
				},
			},
		},
		AvailabilityExceptions: []dtos.AvailabilityExceptionDTO{
			{
				DateTimeRange: dtos.DateTimeRangeDTO{
					Start:    july25,
					Duration: duration1Week,
				},
			},
		},
	}
)

var testCases []CreateAvailabilityScheduleCases = []CreateAvailabilityScheduleCases{
	{
		Info:                  "professor should create schedule",
		CreateScheduleRequest: baseSchedule,
		Professor:             schedule_domain.NewProfessor(professorUuid, "big", "brother", schedule_domain.Schedule{}, []schedule_domain.Booking{}),
		ProfessorBookings: []schedule_domain.Booking{
			schedule_domain.NewBooking(
				professorDate,
				duration1h,
				otherStudentUuid,
			),
		},
		ExpectedResponse: dtos_responses.CreateScheduleResponse{
			WeeklyAvailabilities: []dtos.WeeklyAvailabilityDTO{
				{
					Day: time.Monday,
					TimeRange: dtos.TimeRangeDTO{
						LocalTimeStart: dtos.LocalTimeDTO{
							Hour:   9,
							Minute: 0,
						},
						Duration: utils_time.MustParseDuration("4h"),
					},
				},
			},
			AvailabilityExceptions: []dtos.AvailabilityExceptionDTO{
				{
					DateTimeRange: dtos.DateTimeRangeDTO{
						Start:    july25,
						Duration: duration1Week,
					},
				},
			},
		},
	},
}

func TestCreateSchedule(t *testing.T) {
	for _, cs := range testCases {
		t.Run(cs.Info, func(t *testing.T) {
			professors := schedule_infra.NewInMemProfRepo(cs.Professor, false)
			scheduleSrv := schedule_application.NewScheduleSrv(professors)
			ctx := context.Background()
			ctx = context.WithValue(ctx, auth_jwt.UserId, professorUuid)
			res, err := scheduleSrv.CreateSchedule(ctx, cs.CreateScheduleRequest)
			fmt.Println("isError", cs.IsError)
			if cs.IsError {
				fmt.Println("err", err)
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, schedule_domain.ErrProfessorNotSchedulable)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
				// utils_display.PrettyDisplay("Expected", cs.ExpectedResponse)
				// utils_display.PrettyDisplay("RES", res)
			}
		})
	}
}
