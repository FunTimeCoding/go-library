package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/timeline"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func timelineRow(e *timeline.Entry) gomponents.Node {
	t, f := time.Parse(time.RFC3339, e.Timestamp)
	display := e.Timestamp

	if f == nil {
		display = t.Local().Format("15:04")
	}

	return html.Tr(
		html.Td(gomponents.Text(display)),
		html.Td(gomponents.Text(timeline.FormatDescription(e))),
	)
}
