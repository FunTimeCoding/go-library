package web

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func statusBadge(status string) g.Node {
	return h.Span(
		c.Classes{"badge": true, "badge-" + status: true},
		g.Text(status),
	)
}
