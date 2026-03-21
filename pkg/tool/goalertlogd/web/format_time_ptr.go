package web

import "time"

func formatTimePtr(t *time.Time) string {
	if t == nil {
		return "-"
	}

	return formatTime(*t)
}
