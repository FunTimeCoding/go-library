package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func navigationLink(link, label, currentPath string) g.Node {
	attrs := []g.Node{
		g.Attr("href", link),
		g.Text(label),
	}

	if currentPath == link {
		attrs = append(attrs, g.Attr("aria-current", "page"))
	}

	return h.Li(h.A(attrs...))
}
