package web

import (
	"fmt"
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
	"time"
)

func severityBadge(severity string) g.Node {
	return h.Span(
		c.Classes{"badge": true, "badge-" + severity: true},
		g.Text(severity),
	)
}

func statusBadge(status string) g.Node {
	return h.Span(
		c.Classes{"badge": true, "badge-" + status: true},
		g.Text(status),
	)
}

func formatDuration(d time.Duration) string {
	if d == 0 {
		return "-"
	}

	if d < time.Minute {
		return d.Round(time.Second).String()
	}

	if d < time.Hour {
		return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
	}

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60

	return fmt.Sprintf("%dh%dm", hours, minutes)
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func formatTimePtr(t *time.Time) string {
	if t == nil {
		return "-"
	}

	return formatTime(*t)
}
