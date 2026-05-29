package model_context

import "time"

func parseDurationOrTime(s string) (time.Time, error) {
	d, e := time.ParseDuration(s)

	if e == nil {
		return time.Now().Add(-d), nil
	}

	return time.Parse(time.RFC3339, s)
}
