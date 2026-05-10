package utils_time

import "time"

func MustParseDuration(duration string) time.Duration {
	d, _ := time.ParseDuration(duration)
	return d
}
