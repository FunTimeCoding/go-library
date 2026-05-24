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

	var head []gomponents.Node
	head = append(
		head,
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
		head = append(head, html.Script(html.Src(s)))
	}

	if p.liveEndpoint != "" {
		head = append(
			head,
			html.Script(html.Src(constant.ServerSide)),
		)
	}

	if p.theme != "" {
		head = append(
			head,
			html.StyleEl(gomponents.Raw(p.theme)),
		)
	}

	if p.style != "" {
		head = append(
			head,
			html.StyleEl(gomponents.Raw(p.style)),
		)
	}

	var body []gomponents.Node
	brand := p.identity.Name()

	if brand != "" || p.brandNode != nil {
		node := p.brandNode

		if node == nil {
			node = html.A(
				gomponents.Attr("href", "/"),
				html.Strong(gomponents.Text(brand)),
			)
		}

		var navigation []gomponents.Node

		for _, item := range p.items {
			navigation = append(navigation, item.Render(p.path))
		}

		navigation = append(navigation, p.navigation...)
		body = append(
			body,
			html.Nav(
				html.Class("container"),
				html.Ul(html.Li(node)),
				html.Ul(navigation...),
			),
		)
	}

	var mainAttrs []gomponents.Node
	mainAttrs = append(mainAttrs, html.Class("container"))

	if p.liveEndpoint != "" {
		mainAttrs = append(
			mainAttrs,
			gomponents.Attr("hx-ext", "sse"),
			gomponents.Attr("sse-connect", p.liveEndpoint),
		)
	}

	if len(p.summary) > 0 {
		swap := summaryStrip(p.summary)

		if p.liveEndpoint != "" {
			swap = summaryStripLive(p.summary)
		}

		mainAttrs = append(mainAttrs, swap)
	}

	mainAttrs = append(mainAttrs, gomponents.Group(p.content))
	body = append(body, html.Main(mainAttrs...))
	body = append(body, p.footer...)

	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			gomponents.Attr("data-theme", "dark"),
			html.Head(head...),
			html.Body(body...),
		),
	)
}
