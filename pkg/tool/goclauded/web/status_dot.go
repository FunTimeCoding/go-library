package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func statusDot(lastSeen time.Time) gomponents.Node {
	class := "status-dot status-active"

	if time.Since(lastSeen) > 45*time.Minute {
		class = "status-dot status-stale"
	}

	return html.Span(html.Class(class))
}
