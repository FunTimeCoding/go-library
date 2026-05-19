package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func historyRow(e *event.Event) gomponents.Node {
	return html.Tr(
		html.Td(gomponents.Text(e.CreatedAt.Format("15:04"))),
		html.Td(gomponents.Textf("#%d", e.Identifier)),
		html.Td(gomponents.Text(formatEvent(e))),
	)
}
