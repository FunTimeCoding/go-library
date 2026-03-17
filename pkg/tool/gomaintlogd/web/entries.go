package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) handleEntries(
	w http.ResponseWriter,
	r *http.Request,
) {
	f := &store.Filter{}

	if v := r.URL.Query().Get("system"); v != "" {
		f.System = v
	}

	if v := r.URL.Query().Get("service"); v != "" {
		f.Service = v
	}

	if v := r.URL.Query().Get("user"); v != "" {
		f.User = v
	}

	if v := r.URL.Query().Get("since"); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			f.Since = t
		}
	}

	if v := r.URL.Query().Get("until"); v != "" {
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
			"Entries", "/entries",
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
							h.Name("system"),
							h.Value(f.System),
							h.Placeholder("e.g. worker1"),
						),
					),
					h.Label(
						g.Text("Service"),
						h.Input(
							h.Type("text"),
							h.Name("service"),
							h.Value(f.Service),
							h.Placeholder("e.g. nginx"),
						),
					),
					h.Label(
						g.Text("User"),
						h.Input(
							h.Type("text"),
							h.Name("user"),
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
							h.Name("since"),
						),
					),
					h.Label(
						g.Text("Until"),
						h.Input(
							h.Type("datetime-local"),
							h.Name("until"),
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

func entriesTable(entries []store.Entry) g.Node {
	if len(entries) == 0 {
		return h.P(h.Em(g.Text("No entries found.")))
	}

	var rows []g.Node

	for _, e := range entries {
		target := fmt.Sprintf("detail-%d", e.ID)
		rows = append(
			rows,
			h.Tr(
				h.Class("clickable-row"),
				hx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
				hx.Target("#"+target),
				hx.Swap("outerHTML"),
				h.Td(g.Text(formatTime(e.Timestamp))),
				h.Td(g.Text(e.Action)),
				h.Td(g.Text(e.User)),
				h.Td(g.Text(e.System)),
				h.Td(g.Text(e.Service)),
				h.Td(g.Text(truncate(e.Description, 80))),
			),
			h.Tr(
				h.ID(target),
				h.Style("display:none"),
			),
		)
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Timestamp")),
				h.Th(g.Text("Action")),
				h.Th(g.Text("User")),
				h.Th(g.Text("System")),
				h.Th(g.Text("Service")),
				h.Th(g.Text("Description")),
			),
		),
		h.TBody(g.Group(rows)),
	)
}
