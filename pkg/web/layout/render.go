package layout

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Render() gomponents.Node {
	title := p.identity.Name()

	if p.title != "" {
		title = fmt.Sprintf("%s - %s", title, p.title)
	}

	var headChildren []gomponents.Node
	headChildren = append(
		headChildren,
		html.Meta(html.Charset("utf-8")),
		html.Meta(
			html.Name("viewport"),
			html.Content("width=device-width, initial-scale=1"),
		),
		html.TitleEl(gomponents.Text(title)),
		html.Link(html.Rel("stylesheet"), html.Href(constant.Pico)),
		html.Script(html.Src(constant.Extended)),
	)

	for _, s := range p.script {
		headChildren = append(
			headChildren,
			html.Script(html.Src(s)),
		)
	}

	if p.theme != "" {
		headChildren = append(
			headChildren,
			html.StyleEl(gomponents.Raw(p.theme)),
		)
	}

	if p.style != "" {
		headChildren = append(
			headChildren,
			html.StyleEl(gomponents.Raw(p.style)),
		)
	}

	var bodyChildren []gomponents.Node
	brand := p.identity.Name()

	if brand != "" || p.brandNode != nil {
		node := p.brandNode

		if node == nil {
			node = html.A(
				gomponents.Attr("href", "/"),
				html.Strong(gomponents.Text(brand)),
			)
		}

		var navigationLinks []gomponents.Node

		for _, item := range p.items {
			navigationLinks = append(navigationLinks, item.Render(p.path))
		}

		navigationLinks = append(navigationLinks, p.navigation...)
		bodyChildren = append(
			bodyChildren,
			html.Nav(
				html.Class("container"),
				html.Ul(html.Li(node)),
				html.Ul(navigationLinks...),
			),
		)
	}

	bodyChildren = append(
		bodyChildren,
		html.Main(
			html.Class("container"),
			gomponents.Group(p.content),
		),
	)
	bodyChildren = append(bodyChildren, p.footer...)

	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			gomponents.Attr("data-theme", "dark"),
			html.Head(headChildren...),
			html.Body(bodyChildren...),
		),
	)
}
