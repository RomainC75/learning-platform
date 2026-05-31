package shared_infra

import (
	"time"
)

type DeterministicTimeGenerator struct {
	expectedTime time.Time
	sleepTick    chan struct{}
}

func NewDeterministicTimeGenerator(expectedTime time.Time) *DeterministicTimeGenerator {
	sleepTick := make(chan struct{})
	return &DeterministicTimeGenerator{
		expectedTime: expectedTime,
		sleepTick:    sleepTick,
	}
}

func (dtg *DeterministicTimeGenerator) Now() time.Time {
	return dtg.expectedTime
}
