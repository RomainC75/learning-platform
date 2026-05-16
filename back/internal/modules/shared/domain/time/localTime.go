package shared_domain_time

import (
	"errors"
)

var (
	ErrMoreThan24Hours   = "hours should be positive or zero and no more that 24 hours in a day"
	ErrMoreThan60Minutes = "hours should be positive or zero and no more that 60 minutes in an hour"
)

type LocalTime struct {
	hour   int
	minute int
}

type LocalTimeSnapshot struct {
	Hour   int
	Minute int
}

func NewLocalTime24(hour int, min int) (LocalTime, error) {
	if hour >= 24 || hour < 0 {
		return LocalTime{}, errors.New(ErrMoreThan24Hours)
	} else if min >= 60 || min < 0 {
		return LocalTime{}, errors.New(ErrMoreThan24Hours)
	}
	return LocalTime{
		hour:   hour,
		minute: min,
	}, nil
}

func MustLocalTime24(hour int, min int) LocalTime {
	localTime, err := NewLocalTime24(hour, min)
	if err != nil {
		panic(err)
	}
	return localTime
}

func (lt LocalTime) ToSnapshot() LocalTimeSnapshot {
	return LocalTimeSnapshot{
		Hour:   lt.hour,
		Minute: lt.minute,
	}
}
