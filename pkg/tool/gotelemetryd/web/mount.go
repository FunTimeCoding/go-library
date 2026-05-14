package web

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.heatmap)
	m.HandleFunc("GET /events", s.events)
}
