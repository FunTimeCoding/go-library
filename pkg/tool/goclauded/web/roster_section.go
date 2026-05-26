package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) rosterSection() gomponents.Node {
	sessions := s.service.ListSessions()

	if len(sessions) == 0 {
		return html.P(gomponents.Text("No active sessions."))
	}

	var cards []gomponents.Node

	for i := range sessions {
		labels := s.service.LabelsBySession(sessions[i].Identifier)
		latest := s.service.LatestPulse(sessions[i].Identifier)
		cards = append(
			cards,
			sessionCard(&sessions[i], sessions[i].Lines, labels, latest),
		)
	}

	return html.Div(
		html.Class("roster-grid"),
		gomponents.Group(cards),
	)
}
