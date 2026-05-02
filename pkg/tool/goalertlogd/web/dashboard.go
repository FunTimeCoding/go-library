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
		renderFragment(w, s.topTable())

		return
	}

	lastPoll := s.worker.LastPoll()
	lastPollText := "never"

	if !lastPoll.IsZero() {
		lastPollText = formatTime(lastPoll)
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
					html.Header(gomponents.Text("Total Records")),
					html.P(gomponents.Textf("%d", s.store.MustCount())),
				),
				html.Article(
					html.Header(gomponents.Text("Currently Firing")),
					html.P(gomponents.Textf("%d", s.worker.FiringCount())),
				),
				html.Article(
					html.Header(gomponents.Text("Last Poll")),
					html.P(gomponents.Text(lastPollText)),
				),
			),
			html.H2(gomponents.Text("Top 25 Noisy Alerts (Last 7 Days)")),
			html.Div(
				html.ID("top-table"),
				htmx.Get("/"),
				htmx.Trigger("every 60s"),
				htmx.Swap("innerHTML"),
				s.topTable(),
			),
		),
	)
}
