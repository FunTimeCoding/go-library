package web

import "time"

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}

	return s[:n-3] + "..."
}
