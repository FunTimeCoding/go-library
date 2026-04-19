package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) players(
	w http.ResponseWriter,
	r *http.Request,
) {
	since := time.Now().AddDate(0, -1, 0)
	rows := s.store.PlayerAttendance(since)
	renderPage(
		w,
		layout(
			"Players",
			"/players",
			h.H1(g.Textf("Player Attendance (%d)", len(rows))),
			playersTable(rows),
		),
	)
}
