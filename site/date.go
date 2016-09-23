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
	year := strconv.Itoa(time.Year())
	hours := strconv.Itoa(time.Hour())
	minutes := strconv.Itoa(time.Minute())
	return hours + ":" + minutes + ":" + year
}
