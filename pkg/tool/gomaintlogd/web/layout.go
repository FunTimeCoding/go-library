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
				h.TitleEl(g.Textf("Maintenance Log - %s", title)),
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
						h.Li(h.Strong(g.Text("Maintenance Log"))),
					),
					h.Ul(
						navLink("/", "Dashboard", currentPath),
						navLink("/entries", "Entries", currentPath),
						navLink("/add", "Add Entry", currentPath),
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
