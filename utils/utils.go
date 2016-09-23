package utils

import (
	"fmt"
	"time"
)

func TimeDurationInSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()
}

func TimeDurationInNanoseconds(t time.Time) int64 {
	return time.Since(t).Nanoseconds()
}

func TimeDurationInMilliseconds(t time.Time) float64 {
	return TimeDurationInSeconds(t) * 1000
}

func Check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
