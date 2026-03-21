package web

import (
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) dashboard(
	w http.ResponseWriter,
	r *http.Request,
) {
	if s.isHTMX(r) {
		renderFragment(w, s.recentTable())

		return
	}

	renderPage(
		w,
		layout(
			"Dashboard", "/",
			h.H1(g.Text("Dashboard")),
			h.Div(
				h.Class("summary-cards"),
				h.Article(
					h.Header(g.Text("Total Entries")),
					h.P(g.Textf("%d", s.store.Count())),
				),
			),
			h.H2(g.Text("Recent Entries")),
			h.Div(
				h.ID("recent-table"),
				hx.Get("/"),
				hx.Trigger("every 60s"),
				hx.Swap("innerHTML"),
				s.recentTable(),
			),
		),
	)
}
