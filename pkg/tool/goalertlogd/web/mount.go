package web

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc("GET /recent", s.recent)
	m.HandleFunc("GET /alerts", s.alerts)
}
