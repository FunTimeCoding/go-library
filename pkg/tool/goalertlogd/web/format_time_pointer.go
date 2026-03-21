package web

import "time"

func formatTimePointer(t *time.Time) string {
	if t == nil {
		return "-"
	}

	return formatTime(*t)
}
