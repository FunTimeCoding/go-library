package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func layout(
	currentPath string,
	content ...gomponents.Node,
) gomponents.Node {
	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(
					html.Name("viewport"),
					html.Content("width=device-width, initial-scale=1"),
				),
				html.TitleEl(gomponents.Text(constant.Identity.Name())),
				html.Link(
					html.Rel("stylesheet"),
					html.Href(
						"https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css",
					),
				),
				html.Script(
					html.Src("https://unpkg.com/htmx.org@2.0.4"),
				),
			),
			html.Body(
				html.Nav(
					html.Class("container"),
					html.Ul(
						html.Li(
							html.Strong(
								gomponents.Text(constant.Identity.Name()),
							),
						),
					),
					html.Ul(
						navigationLink("/", "Heatmap", currentPath),
						navigationLink("/events", "Events", currentPath),
					),
				),
				html.Main(
					html.Class("container"),
					gomponents.Group(content),
				),
			),
		),
	)
}
