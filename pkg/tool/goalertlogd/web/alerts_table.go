package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func alertsTable(records []store.Record) g.Node {
	if len(records) == 0 {
		return h.P(h.Em(g.Text("No occurrences found.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Fingerprint")),
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
						h.Td(g.Text(r.Fingerprint)),
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
