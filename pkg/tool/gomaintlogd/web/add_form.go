package web

import (
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func addForm() g.Node {
	return h.Form(
		hx.Post("/add"),
		hx.Target("#add-form"),
		hx.Swap("outerHTML"),
		h.ID("add-form"),
		h.Label(
			g.Text("Action (required)"),
			h.Input(
				h.Type("text"),
				h.Name("action"),
				h.Required(),
				h.Placeholder("e.g. restarted web server"),
			),
		),
		h.Label(
			g.Text("User (required)"),
			h.Input(
				h.Type("text"),
				h.Name("user"),
				h.Required(),
				h.Placeholder("e.g. jdoe"),
			),
		),
		h.Div(
			h.Class("grid"),
			h.Label(
				g.Text("System"),
				h.Input(
					h.Type("text"),
					h.Name("system"),
					h.Placeholder("e.g. worker1"),
				),
			),
			h.Label(
				g.Text("Service"),
				h.Input(
					h.Type("text"),
					h.Name("service"),
					h.Placeholder("e.g. nginx"),
				),
			),
		),
		h.Label(
			g.Text("Description"),
			h.Textarea(
				h.Name("description"),
				h.Placeholder("Optional details..."),
				h.Rows("3"),
			),
		),
		h.Button(h.Type("submit"), g.Text("Add Entry")),
	)
}
