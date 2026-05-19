package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) rosterSection() gomponents.Node {
	sessions := s.service.Store.ListSessions()

	if len(sessions) == 0 {
		return html.P(gomponents.Text("No active sessions."))
	}

	var cards []gomponents.Node

	for _, e := range sessions {
		cards = append(cards, sessionCard(&e))
	}

	return html.Div(
		html.Class("roster-grid"),
		gomponents.Group(cards),
	)
}
