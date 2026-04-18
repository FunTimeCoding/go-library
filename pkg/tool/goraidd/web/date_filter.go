package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func dateFilter(startValue, endValue string) g.Node {
	return h.Form(
		h.Class("filter-form"),
		h.Method("get"),
		h.Action("/"),
		h.Div(
			h.Class("grid"),
			h.Label(
				g.Text("From"),
				h.Input(
					h.Type("datetime-local"),
					h.Name("start"),
					h.Value(startValue),
				),
			),
			h.Label(
				g.Text("To"),
				h.Input(
					h.Type("datetime-local"),
					h.Name("end"),
					h.Value(endValue),
				),
			),
			h.Button(
				h.Type("submit"),
				g.Text("Filter"),
			),
		),
	)
}
