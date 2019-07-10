package main

import (
	"fmt"
)

const (
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 24
)

func resolveTime(seconds int) (day int, hour int, minutes int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minutes = seconds / SecondsPerMinute
	return day, hour, minutes
}

func main() {
	seconds := 1000
	day, hour, minutes := resolveTime(seconds)
	fmt.Println(day, hour, minutes)
}
