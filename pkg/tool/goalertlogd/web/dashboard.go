package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) handleDashboard(
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

func (s *Server) topTable() g.Node {
	now := time.Now()
	records := s.store.Top(25, now.Add(-7*24*time.Hour), now)

	if len(records) == 0 {
		return h.P(h.Em(g.Text("No alerts in the last 7 days.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Name")),
				h.Th(g.Text("Count")),
				h.Th(g.Text("Avg Duration")),
				h.Th(g.Text("Firing")),
				h.Th(g.Text("Severity")),
			),
		),
		h.TBody(
			g.Map(
				records, func(r store.TopRecord) g.Node {
					return h.Tr(
						h.Td(
							h.A(
								h.Class("alert-link"),
								h.Href(
									fmt.Sprintf(
										"/alerts?name=%s",
										r.Name,
									),
								),
								g.Text(r.Name),
							),
						),
						h.Td(g.Textf("%d", r.Count)),
						h.Td(g.Text(formatDuration(r.AverageDuration))),
						h.Td(g.Textf("%d", r.CurrentlyFiring)),
						h.Td(severityBadge(r.Severity)),
					)
				},
			),
		),
	)
}
