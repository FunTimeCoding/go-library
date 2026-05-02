package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
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
			html.H1(gomponents.Textf("Raids (%d)", len(rows))),
			raidsTable(rows),
		),
	)
}
