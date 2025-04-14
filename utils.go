package main

import (
	"strconv"
	"time"
)

func remainingTimeToString(rt time.Duration) (string, string) {
	minutes := strconv.Itoa(int(rt.Seconds() / 60))
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}

	seconds := strconv.Itoa(int(rt.Seconds()) % 60)
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}

	return minutes, seconds
}

func secondsToTimeString(s int) string {
	str := ""
	minutes := strconv.Itoa(s / 60)
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	str += minutes
	str += ":"
	seconds := strconv.Itoa(s % 60)
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}
	str += seconds

	return str
}
