package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
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
			"Dashboard",
			"/",
			html.H1(gomponents.Text("Dashboard")),
			html.Div(
				html.Class("summary-cards"),
				html.Article(
					html.Header(gomponents.Text("Total Entries")),
					html.P(gomponents.Textf("%d", s.store.Count())),
				),
			),
			html.H2(gomponents.Text("Recent Entries")),
			html.Div(
				html.ID("recent-table"),
				htmx.Get("/"),
				htmx.Trigger("every 60s"),
				htmx.Swap("innerHTML"),
				s.recentTable(),
			),
		),
	)
}
