package web

import "time"

func sinceFromWindow(window string) string {
	d, e := time.ParseDuration(window)

	if e != nil {
		d = 24 * time.Hour
	}

	return time.Now().UTC().Add(-d).Format(time.RFC3339)
}
