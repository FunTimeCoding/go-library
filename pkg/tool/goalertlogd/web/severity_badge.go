package web

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func severityBadge(severity string) g.Node {
	return h.Span(
		c.Classes{"badge": true, "badge-" + severity: true},
		g.Text(severity),
	)
}
