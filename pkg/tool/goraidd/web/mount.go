package web

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.logs)
	m.HandleFunc("POST /generate", s.generate)
	m.HandleFunc("GET /reports", s.reports)
	m.HandleFunc("GET /reports/{fileName}", s.reportDownload)
	m.HandleFunc("POST /reports/{fileName}/delete", s.reportDelete)
}
