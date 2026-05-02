package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) alerts(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "missing name parameter", http.StatusBadRequest)

		return
	}

	records := s.store.MustByName(name)
	renderPage(
		w,
		layout(
			name,
			"",
			html.H1(gomponents.Textf("Alert: %s", name)),
			html.P(gomponents.Textf("%d occurrences", len(records))),
			alertsTable(records),
		),
	)
}
