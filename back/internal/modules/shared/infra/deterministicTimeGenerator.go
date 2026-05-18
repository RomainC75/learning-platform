package shared_infra

import (
	"time"
)

type DeterministicTimeGenerator struct {
	expectedTime time.Time
}

func NewDeterministicTimeGenerator(expectedTime time.Time) *DeterministicTimeGenerator {
	return &DeterministicTimeGenerator{
		expectedTime: expectedTime,
	}
}

func (dtg *DeterministicTimeGenerator) Now() time.Time {
	return dtg.expectedTime
}
