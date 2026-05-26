package layout

import (
	"maragu.dev/gomponents"
	"strings"
)

func SummaryStripContent(items []string) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			gomponents.Attr(
				"style",
				"padding-top:0;padding-bottom:0.5rem;opacity:0.6;font-size:0.85rem",
			),
			gomponents.Text(strings.Join(items, " · ")),
		},
	)
}
