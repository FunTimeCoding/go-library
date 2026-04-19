package web

import "fmt"

func perSecond(
	total int,
	seconds float64,
) string {
	if seconds <= 0 {
		return "—"
	}

	return fmt.Sprintf("%.1f", float64(total)/seconds)
}
