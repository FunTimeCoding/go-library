package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) entries(
	w http.ResponseWriter,
	r *http.Request,
) {
	f := store.NewFilter()

	if v := r.URL.Query().Get(constant.System); v != "" {
		f.System = v
	}

	if v := r.URL.Query().Get(constant.Service); v != "" {
		f.Service = v
	}

	if v := r.URL.Query().Get(constant.User); v != "" {
		f.User = v
	}

	if v := r.URL.Query().Get(constant.Since); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			f.Since = t
		}
	}

	if v := r.URL.Query().Get(constant.Until); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			f.Until = t
		}
	}

	entries, e := s.store.List(f)
	errors.PanicOnError(e)
	table := entriesTable(entries)

	if s.isHTMX(r) {
		renderFragment(w, table)

		return
	}

	renderPage(
		w,
		layout(
			"Entries",
			"/entries",
			h.H1(g.Text("Entries")),
			h.Form(
				h.Class("filter-form"),
				hx.Get("/entries"),
				hx.Target("#entries-table"),
				hx.Swap("innerHTML"),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("System"),
						h.Input(
							h.Type("text"),
							h.Name(constant.System),
							h.Value(f.System),
							h.Placeholder("e.g. worker1"),
						),
					),
					h.Label(
						g.Text("Service"),
						h.Input(
							h.Type("text"),
							h.Name(constant.Service),
							h.Value(f.Service),
							h.Placeholder("e.g. nginx"),
						),
					),
					h.Label(
						g.Text("User"),
						h.Input(
							h.Type("text"),
							h.Name(constant.User),
							h.Value(f.User),
						),
					),
				),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("Since"),
						h.Input(
							h.Type("datetime-local"),
							h.Name(constant.Since),
						),
					),
					h.Label(
						g.Text("Until"),
						h.Input(
							h.Type("datetime-local"),
							h.Name(constant.Until),
						),
					),
					h.Label(
						g.Raw("&nbsp;"),
						h.Button(h.Type("submit"), g.Text("Filter")),
					),
				),
			),
			h.Div(h.ID("entries-table"), table),
		),
	)
}
