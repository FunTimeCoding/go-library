package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) playerDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	account := r.PathValue("account")
	rows := s.store.PlayerDetail(account)
	renderPage(
		w,
		layout(
			account,
			"/players",
			h.H1(g.Text(account)),
			playerDetailTable(rows),
		),
	)
}
