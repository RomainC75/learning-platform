package schedule_unit

import (
	"fmt"
	"language-learning/internal/api/dtos"
	dtos_requests "language-learning/internal/api/dtos/request"
	dtos_responses "language-learning/internal/api/dtos/responses"
	schedule_application "language-learning/internal/modules/schedule/application"
	schedule_domain "language-learning/internal/modules/schedule/domain"
	utils_display "language-learning/utils/display"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CreateAvailabilityScheduleCases struct {
	Info                  string
	CreateScheduleRequest dtos_requests.CreateScheduleRequest
	ExpectedResponse      dtos_responses.CreateScheduleResponse
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
			testDriver.SetSchedule(schedule_domain.Schedule{})
			res, err := testDriver.CreateSchedule(cs.CreateScheduleRequest)

			fmt.Println("isError", cs.IsError)
			if cs.IsError {
				assert.NotNil(t, err, cs.ErrorMessage)
				assert.EqualError(t, err, cs.ErrorMessage)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, cs.ExpectedResponse, res)
				// test profesor schedule is saved
				scheduleTemoin := schedule_application.NewScheduleFromDto(cs.CreateScheduleRequest)

				utils_display.PrettyDisplay("schedule --->", scheduleTemoin.ToSnapshot())
				utils_display.PrettyDisplay("Saved one  --->", testDriver.GetSavedSchedule().ToSnapshot())
				assert.Equal(t, scheduleTemoin.ToSnapshot(), testDriver.GetSavedSchedule().ToSnapshot())
			}
		})
	}
}
