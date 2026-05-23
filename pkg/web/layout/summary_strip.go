package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func summaryStrip(items []string) gomponents.Node {
	return html.Div(
		html.Class("container"),
		gomponents.Attr(
			"style",
			"padding-top:0;padding-bottom:0.5rem;opacity:0.6;font-size:0.85rem",
		),
		gomponents.Text(strings.Join(items, " · ")),
	)
}
