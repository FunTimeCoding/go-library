package time

import "time"

func PublicHoliday(t time.Time) bool {
	for _, d := range PublicHolidays {
		if d.Year() == t.Year() &&
			d.Month() == t.Month() &&
			d.Day() == t.Day() {
			return true
		}
	}

	return false
}
