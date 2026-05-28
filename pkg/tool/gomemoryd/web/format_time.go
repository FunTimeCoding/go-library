package web

import "time"

func formatTime(raw string) string {
	t, e := time.Parse(time.RFC3339, raw)

	if e != nil {
		return raw
	}

	return t.Local().Format("2006-01-02 15:04")
}
