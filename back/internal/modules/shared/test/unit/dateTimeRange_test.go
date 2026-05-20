package unit

import (
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	utils_time "language-learning/utils/time"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type DateTimeRangeTestCases struct {
	Message   string
	Expected  bool
	Start1    time.Time
	Start2    time.Time
	Duration1 time.Duration
	Duration2 time.Duration
}

var tcs []DateTimeRangeTestCases = []DateTimeRangeTestCases{
	{
		Message:   "should overlap 1",
		Expected:  true,
		Start1:    time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 2, 12, 30, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
	{
		Message:   "should overlap 2",
		Expected:  true,
		Start1:    time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
	{
		Message:   "should not overlap 1",
		Expected:  false,
		Start1:    time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 2, 13, 0, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
	{
		Message:   "should not overlap 2",
		Expected:  false,
		Start1:    time.Date(2026, time.April, 2, 13, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
	{
		Message:   "should not overlap, cause different days 	1",
		Expected:  false,
		Start1:    time.Date(2026, time.April, 2, 13, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 1, 13, 0, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
	{
		Message:   "should not overlap, cause different days    2",
		Expected:  false,
		Start1:    time.Date(2026, time.April, 1, 13, 0, 0, 0, time.UTC),
		Start2:    time.Date(2026, time.April, 2, 13, 0, 0, 0, time.UTC),
		Duration1: utils_time.MustParseDuration("1h"),
		Duration2: utils_time.MustParseDuration("1h"),
	},
}

func TestDateTimeRange(t *testing.T) {
	for _, tc := range tcs {
		t.Run(tc.Message, func(t *testing.T) {
			dtr1 := shared_domain_time.NewDateTimeRange(tc.Start1, tc.Duration1)
			dtr2 := shared_domain_time.NewDateTimeRange(tc.Start2, tc.Duration2)
			res := dtr1.IsOverlapWith(dtr2)

			assert.Equal(t, tc.Expected, res)
		})
	}
}
