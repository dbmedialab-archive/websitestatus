package site

import (
	"strconv"
	"time"
)

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
