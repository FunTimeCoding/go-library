package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

func statusBadge(status string) gomponents.Node {
	return html.Span(
		components.Classes{"badge": true, fmt.Sprintf("badge-%s", status): true},
		gomponents.Text(status),
	)
}
