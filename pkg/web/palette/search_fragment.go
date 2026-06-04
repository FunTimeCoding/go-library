package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SearchFragment(
	endpoint string,
	placeholder string,
) gomponents.Node {
	return html.Div(
		html.Class("palette-body"),
		html.Input(
			html.ID("palette-search-input"),
			html.Type("text"),
			html.Class("palette-input"),
			gomponents.Attr("placeholder", placeholder),
			gomponents.Attr("autocomplete", "off"),
			gomponents.Attr("hx-get", endpoint),
			gomponents.Attr("hx-trigger", "keyup changed delay:200ms"),
			gomponents.Attr("hx-target", "#palette-search-results"),
			gomponents.Attr("hx-swap", "outerHTML"),
			gomponents.Attr("name", "q"),
			gomponents.Attr("autofocus", ""),
		),
		html.Div(html.ID("palette-search-results")),
	)
}
