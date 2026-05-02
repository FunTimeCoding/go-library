package web

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func dateFilter(startValue, endValue string) gomponents.Node {
	return html.Form(
		html.Class("filter-form"),
		html.Method(constant.Get),
		html.Action("/"),
		html.Div(
			html.Class("grid"),
			html.Label(
				gomponents.Text("From"),
				html.Input(
					html.Type("datetime-local"),
					html.Name("start"),
					html.Value(startValue),
				),
			),
			html.Label(
				gomponents.Text("To"),
				html.Input(
					html.Type("datetime-local"),
					html.Name("end"),
					html.Value(endValue),
				),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text("Filter"),
			),
		),
	)
}
