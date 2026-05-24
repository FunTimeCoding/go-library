package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func summaryStripLive(items []string) gomponents.Node {
	return html.Div(
		gomponents.Attr("sse-swap", SummaryStrip),
		SummaryStripContent(items),
	)
}
