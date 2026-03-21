package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func layout(
	title string,
	currentPath string,
	content ...g.Node,
) g.Node {
	return h.Doctype(
		h.HTML(
			h.Lang("en"),
			h.Head(
				h.Meta(h.Charset("utf-8")),
				h.Meta(
					h.Name("viewport"),
					h.Content("width=device-width, initial-scale=1"),
				),
				h.TitleEl(g.Textf("Alert Log - %s", title)),
				h.Link(
					h.Rel("stylesheet"),
					h.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"),
				),
				h.Script(h.Src("https://unpkg.com/htmx.org@2.0.4")),
				h.StyleEl(g.Raw(inlineCSS)),
			),
			h.Body(
				h.Nav(
					h.Class("container"),
					h.Ul(
						h.Li(h.Strong(g.Text("Alert Log"))),
					),
					h.Ul(
						navLink("/", "Dashboard", currentPath),
						navLink("/recent", "Recent", currentPath),
					),
				),
				h.Main(
					h.Class("container"),
					g.Group(content),
				),
			),
		),
	)
}

func navLink(href, label, currentPath string) g.Node {
	attrs := []g.Node{
		g.Attr("href", href),
		g.Text(label),
	}

	if currentPath == href {
		attrs = append(attrs, g.Attr("aria-current", "page"))
	}

	return h.Li(h.A(attrs...))
}

const inlineCSS = `
.badge { padding: 2px 8px; border-radius: 4px; font-size: 0.8em; font-weight: bold; display: inline-block; }
.badge-critical { background: #f8d7da; color: #721c24; }
.badge-warning { background: #fff3cd; color: #856404; }
.badge-opsgenie, .badge-adition-opsgenie { background: #cce5ff; color: #004085; }
.badge-info { background: #d1ecf1; color: #0c5460; }
.badge-firing { background: #f8d7da; color: #721c24; }
.badge-resolved { background: #d4edda; color: #155724; }
.summary-cards { display: flex; gap: 1rem; margin-bottom: 2rem; flex-wrap: wrap; }
.summary-cards article { flex: 1; min-width: 150px; text-align: center; }
table { margin-top: 1rem; }
.alert-link { text-decoration: none; }
.alert-link:hover { text-decoration: underline; }
.filter-form { margin-bottom: 1rem; }
.filter-form .grid { align-items: end; }
`
