package time

import "time"

func PublicHoliday(t time.Time) bool {
	for _, element := range PublicHolidays {
		if element.Year() == t.Year() &&
			element.Month() == t.Month() &&
			element.Day() == t.Day() {
			return true
		}
	}

	return false
}
