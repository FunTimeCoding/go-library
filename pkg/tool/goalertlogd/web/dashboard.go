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
		renderFragment(w, s.topTable())

		return
	}

	lastPoll := s.poller.LastPoll()
	lastPollText := "never"

	if !lastPoll.IsZero() {
		lastPollText = formatTime(lastPoll)
	}

	renderPage(
		w,
		layout(
			"Dashboard", "/",
			h.H1(g.Text("Dashboard")),
			h.Div(
				h.Class("summary-cards"),
				h.Article(
					h.Header(g.Text("Total Records")),
					h.P(g.Textf("%d", s.store.Count())),
				),
				h.Article(
					h.Header(g.Text("Currently Firing")),
					h.P(g.Textf("%d", s.poller.FiringCount())),
				),
				h.Article(
					h.Header(g.Text("Last Poll")),
					h.P(g.Text(lastPollText)),
				),
			),
			h.H2(g.Text("Top 25 Noisy Alerts (Last 7 Days)")),
			h.Div(
				h.ID("top-table"),
				hx.Get("/"),
				hx.Trigger("every 60s"),
				hx.Swap("innerHTML"),
				s.topTable(),
			),
		),
	)
}
