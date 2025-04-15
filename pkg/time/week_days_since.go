package time

import "time"

func WeekDaysSince(
	start time.Time,
	now time.Time,
) int {
	var result int

	if start.After(now) {
		start, now = now, start
	}

	for d := start; d.Before(now); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday {
			continue
		}

		if d.Weekday() == time.Sunday {
			continue
		}

		if PublicHoliday(d) {
			continue
		}

		result++
	}

	return result
}
