package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func filterBar(o *store.QueryOption) gomponents.Node {
	var active []gomponents.Node

	if o.Tool != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("tool=%s ", o.Tool)),
				html.A(
					gomponents.Attr("href", "/events"),
					gomponents.Text("×"),
				),
			),
		)
	}

	if o.Surface != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("surface=%s ", o.Surface)),
				html.A(
					gomponents.Attr("href", "/events"),
					gomponents.Text("×"),
				),
			),
		)
	}

	if o.Actor != "" {
		active = append(
			active,
			html.Span(
				gomponents.Text(fmt.Sprintf("actor=%s ", o.Actor)),
				html.A(
					gomponents.Attr("href", "/events"),
					gomponents.Text("×"),
				),
			),
		)
	}

	if len(active) == 0 {
		return gomponents.Text("")
	}

	return html.P(
		gomponents.Text("Filtered: "),
		gomponents.Group(active),
	)
}
