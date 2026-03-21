package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func navigationLink(href, label, currentPath string) g.Node {
	attrs := []g.Node{
		g.Attr("href", href),
		g.Text(label),
	}

	if currentPath == href {
		attrs = append(attrs, g.Attr("aria-current", "page"))
	}

	return h.Li(h.A(attrs...))
}
