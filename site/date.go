package site

import (
	"strconv"
	"time"
)

type Date struct {
	Date time.Time `json: date`
	Day  string    `json: day`
	Time string    `json: time`
}

func DateToString(time time.Time) string {
	hours := strconv.Itoa(time.Hour())
	if time.Hour() < 10 {
		hours = "0" + strconv.Itoa(time.Hour())
	}
	minutes := strconv.Itoa(time.Minute())
	return hours + ":" + minutes
}
