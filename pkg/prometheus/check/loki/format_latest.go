package loki

import "github.com/funtimecoding/go-library/pkg/time"

func formatLatest(e *overview) string {
	if e.Latest.IsZero() {
		return ""
	}

	return e.Latest.Format(time.DateMinute)
}
