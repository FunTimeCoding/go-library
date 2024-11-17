package log

import "time"

func DaysContains(
	days []time.Time,
	day time.Time,
) bool {
	for _, d := range days {
		if d.Year() == day.Year() && d.YearDay() == day.YearDay() {
			return true
		}
	}

	return false
}
