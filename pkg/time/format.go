package time

import "time"

func Format(t time.Time) string {
	return t.Format(DateMinute)
}
