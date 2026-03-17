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
				h.TitleEl(g.Textf("Maintenance Log — %s", title)),
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
.summary-cards { display: flex; gap: 1rem; margin-bottom: 2rem; flex-wrap: wrap; }
.summary-cards article { flex: 1; min-width: 150px; text-align: center; }
table { margin-top: 1rem; }
.filter-form { margin-bottom: 1rem; }
.filter-form .grid { align-items: end; }
.success { background: #d4edda; color: #155724; padding: 1rem; border-radius: 4px; margin-bottom: 1rem; }
.clickable-row { cursor: pointer; }
.clickable-row:hover { background: var(--pico-table-row-stripped-background-color); }
.detail-row td { padding: 0; }
.detail-content { padding: 1rem; }
.detail-content pre { white-space: pre-wrap; word-break: break-word; margin: 0.5rem 0; }
.detail-actions { display: flex; gap: 0.5rem; margin-top: 1rem; }
.detail-actions button { width: auto; margin-bottom: 0; }
.edit-form { padding: 1rem; }
.edit-form button { width: auto; }
`
