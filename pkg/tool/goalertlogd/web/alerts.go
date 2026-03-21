package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
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

	records := s.store.ByName(name)
	renderPage(
		w,
		layout(
			name, "",
			h.H1(g.Textf("Alert: %s", name)),
			h.P(g.Textf("%d occurrences", len(records))),
			alertsTable(records),
		),
	)
}
