package web

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.logs)
	m.HandleFunc("POST /generate", s.generate)
	m.HandleFunc("GET /reports", s.reports)
	m.HandleFunc("GET /reports/{fileName}", s.reportDownload)
	m.HandleFunc("POST /reports/{fileName}/delete", s.reportDelete)
	m.HandleFunc("GET /players", s.players)
	m.HandleFunc("GET /players/{account}", s.playerDetail)
	m.Handle("GET /static/", http.FileServerFS(staticFiles))
}
