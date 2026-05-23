package schedule_unit

import (
	"language-learning/internal/api/dtos"
	dtos_requests "language-learning/internal/api/dtos/request"
	dtos_responses "language-learning/internal/api/dtos/responses"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CreateAvailabilityScheduleCases struct {
	Info                  string
	CreateScheduleRequest dtos_requests.CreateScheduleRequest
	ExpectedResponse      dtos_responses.CreateScheduleResponse
	ShouldResetSchedule   bool
	IsProfessorMissing    bool
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
	baseScheduleResponse = dtos_responses.CreateScheduleResponse{
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
		ExpectedResponse:      baseScheduleResponse,
	},
	{
		Info:                  "professor should create schedule if there is no schedule",
		CreateScheduleRequest: baseSchedule,
		ShouldResetSchedule:   true,
		ExpectedResponse:      baseScheduleResponse,
	},
	{
		Info:                  "professor should not create schedule if professor is missing",
		CreateScheduleRequest: baseSchedule,
		ExpectedResponse:      dtos_responses.CreateScheduleResponse{},
		IsProfessorMissing:    true,
		IsError:               true,
		ErrorMessage:          schedule_domain.ErrProfessorNotFound,
	},
}

func TestCreateSchedule(t *testing.T) {
	for _, cs := range testCases {
		t.Run(cs.Info, func(t *testing.T) {

			testDriver := NewScheduleTestDriver()
			testDriver.NewScheduleService(cs.IsProfessorMissing)
			if cs.ShouldResetSchedule {
				testDriver.ResetSchedule()
			}
			testDriver.SetSchedule(schedule_domain.Schedule{})
			res, err := testDriver.CreateSchedule(cs.CreateScheduleRequest)

			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, cs.ErrorMessage)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
				scheduleTemoin := schedule_application.NewScheduleFromDto(cs.CreateScheduleRequest)
				assert.Equal(t, scheduleTemoin.ToSnapshot(), testDriver.GetSavedSchedule().ToSnapshot())
			}
		})
	}
}
