package web

import (
	"fmt"
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func severityBadge(severity string) g.Node {
	return h.Span(
		c.Classes{"badge": true, fmt.Sprintf("badge-%s", severity): true},
		g.Text(severity),
	)
}
