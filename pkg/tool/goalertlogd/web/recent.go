package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) handleRecent(
	w http.ResponseWriter,
	r *http.Request,
) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v := r.URL.Query().Get("start"); v != "" {
		if t, e := time.Parse("2006-01-02T15:04", v); e == nil {
			start = t
		}
	}

	if v := r.URL.Query().Get("end"); v != "" {
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
			"Recent", "/recent",
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
							h.Name("start"),
							h.Value(start.Format("2006-01-02T15:04")),
						),
					),
					h.Label(
						g.Text("End"),
						h.Input(
							h.Type("datetime-local"),
							h.Name("end"),
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

func recentTable(records []store.Record) g.Node {
	if len(records) == 0 {
		return h.P(h.Em(g.Text("No alerts in this time range.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Name")),
				h.Th(g.Text("Severity")),
				h.Th(g.Text("Summary")),
				h.Th(g.Text("Start")),
				h.Th(g.Text("End")),
				h.Th(g.Text("Status")),
			),
		),
		h.TBody(
			g.Map(
				records, func(r store.Record) g.Node {
					status := "firing"

					if r.End != nil {
						status = "resolved"
					}

					return h.Tr(
						h.Td(
							h.A(
								h.Class("alert-link"),
								h.Href("/alerts?name="+r.Name),
								g.Text(r.Name),
							),
						),
						h.Td(severityBadge(r.Severity)),
						h.Td(g.Text(truncate(r.Summary, 80))),
						h.Td(g.Text(formatTime(r.Start))),
						h.Td(g.Text(formatTimePtr(r.End))),
						h.Td(statusBadge(status)),
					)
				},
			),
		),
	)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}

	return s[:n-3] + "..."
}
