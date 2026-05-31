package shared_infra

import "time"

type TimeGenerator struct {
}

func NewTimeGenerator() *TimeGenerator {
	return &TimeGenerator{}
}

func (tg *TimeGenerator) Now() time.Time {
	return time.Now()
}

func (tg *TimeGenerator) Sleep(d time.Duration) {
	time.Sleep(d)
}
