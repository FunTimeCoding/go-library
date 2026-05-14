package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func navigationLink(
	path string,
	label string,
	currentPath string,
) gomponents.Node {
	if path == currentPath {
		return html.Li(
			html.A(
				gomponents.Attr("href", path),
				gomponents.Attr("aria-current", "page"),
				gomponents.Text(label),
			),
		)
	}

	return html.Li(
		html.A(
			gomponents.Attr("href", path),
			gomponents.Text(label),
		),
	)
}
