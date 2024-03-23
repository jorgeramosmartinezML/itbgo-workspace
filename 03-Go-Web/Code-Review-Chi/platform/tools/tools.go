package tools

import "time"

func IsValidYear(year int) bool {
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	return date.Year() == year
}
