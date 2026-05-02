package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func navigationLink(link, label, currentPath string) gomponents.Node {
	attrs := []gomponents.Node{
		gomponents.Attr("href", link),
		gomponents.Text(label),
	}

	if currentPath == link {
		attrs = append(attrs, gomponents.Attr("aria-current", "page"))
	}

	return html.Li(html.A(attrs...))
}
