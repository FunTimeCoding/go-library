package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
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
			html.H1(gomponents.Text(account)),
			playerDetailTable(rows),
		),
	)
}
