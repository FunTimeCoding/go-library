package web

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc("GET /event", s.event)
	m.HandleFunc("GET /sessions", s.sessionsPage)
	m.HandleFunc("GET /messages", s.messagesPage)
	m.HandleFunc("GET /history", s.historyPage)
	s.conversations.Mount(m)
}
