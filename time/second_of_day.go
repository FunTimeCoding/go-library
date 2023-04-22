package time

import "time"

func SecondOfDay(now time.Time) int {
	y, m, d := now.Date()

	return int(
		now.Sub(
			time.Date(
				y,
				m,
				d,
				0,
				0,
				0,
				0,
				now.Location(),
			),
		).Seconds(),
	)
}
