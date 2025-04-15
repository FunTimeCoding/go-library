package time

import "time"

func NewDay(
	year int,
	m time.Month,
	day int,
) time.Time {
	return time.Date(year, m, day, 0, 0, 0, 0, time.UTC)
}
