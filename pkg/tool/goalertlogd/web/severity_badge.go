package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

func severityBadge(severity string) gomponents.Node {
	return html.Span(
		components.Classes{"badge": true, fmt.Sprintf("badge-%s", severity): true},
		gomponents.Text(severity),
	)
}
