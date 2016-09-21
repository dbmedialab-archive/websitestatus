package utils

import (
	"fmt"
	"time"
)

func TimeDurationInSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()
}

func Check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
