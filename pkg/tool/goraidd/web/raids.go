package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) raids(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows := s.store.Raids()
	renderPage(
		w,
		layout(
			"Raids",
			"/raids",
			h.H1(g.Textf("Raids (%d)", len(rows))),
			raidsTable(rows),
		),
	)
}
