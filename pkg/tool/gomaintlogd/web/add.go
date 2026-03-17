package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) handleAddForm(
	w http.ResponseWriter,
	_ *http.Request,
) {
	renderPage(
		w,
		layout(
			"Add Entry", "/add",
			h.H1(g.Text("Add Entry")),
			h.Div(h.ID("result")),
			addForm(),
		),
	)
}

func (s *Server) handleAddSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())

	e := &store.Entry{
		Action:      r.FormValue("action"),
		User:        r.FormValue("user"),
		System:      r.FormValue("system"),
		Service:     r.FormValue("service"),
		Description: r.FormValue("description"),
	}

	errors.PanicOnError(s.store.Add(e))

	renderFragment(
		w,
		h.Div(
			h.Div(
				h.Class("success"),
				g.Textf(
					"Entry added: %s by %s",
					e.Action,
					e.User,
				),
			),
			addForm(),
		),
	)
}

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
