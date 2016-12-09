package utils

import (
	"fmt"
	"strconv"
	"time"
)

// TimeDurationInSeconds converts time duration to seconds
func TimeDurationInSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()
}

// TimeDurationInNanoseconds converts time duration to nanoseconds
func TimeDurationInNanoseconds(t time.Time) int64 {
	return time.Since(t).Nanoseconds()
}

// TimeDurationInMilliseconds converts time duration to milliseconds
func TimeDurationInMilliseconds(t time.Time) float64 {
	return TimeDurationInSeconds(t) * 1000
}

// Ok checks if there is an error
func Ok(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// DateToString converts time to string
func DateToString(time time.Time) string {
	hours := strconv.Itoa(time.Hour())
	if time.Hour() < 10 {
		hours = "0" + hours
	}
	minutes := strconv.Itoa(time.Minute())
	if time.Minute() < 10 {
		minutes = "0" + minutes
	}
	seconds := strconv.Itoa(time.Second())
	if time.Second() < 10 {
		seconds = "0" + seconds
	}
	return hours + ":" + minutes + ":" + seconds
}
