package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) recent(
	w http.ResponseWriter,
	r *http.Request,
) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v := r.URL.Query().Get(constant.Start); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			start = t
		}
	}

	if v := r.URL.Query().Get(constant.End); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			end = t
		}
	}

	table := recentTable(s.store.ByTimeRange(start, end))

	if s.isHTMX(r) {
		renderFragment(w, table)

		return
	}

	renderPage(
		w,
		layout(
			"Recent",
			"/recent",
			h.H1(g.Text("Recent Alerts")),
			h.Form(
				h.Class("filter-form"),
				hx.Get("/recent"),
				hx.Target("#recent-table"),
				hx.Swap("innerHTML"),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("Start"),
						h.Input(
							h.Type("datetime-local"),
							h.Name(constant.Start),
							h.Value(start.Format("2006-01-02T15:04")),
						),
					),
					h.Label(
						g.Text("End"),
						h.Input(
							h.Type("datetime-local"),
							h.Name(constant.End),
							h.Value(end.Format("2006-01-02T15:04")),
						),
					),
					h.Label(
						g.Raw("&nbsp;"),
						h.Button(h.Type("submit"), g.Text("Filter")),
					),
				),
			),
			h.Div(h.ID("recent-table"), table),
		),
	)
}
