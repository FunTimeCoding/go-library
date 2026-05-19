package conversations

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /conversations", s.page)
	m.HandleFunc("GET /conversations/sidebar", s.sidebar)
	m.HandleFunc("GET /conversations/{id}/rename", s.renameForm)
	m.HandleFunc("POST /conversations/{id}/rename", s.renameSubmit)
	m.HandleFunc("GET /conversations/{id}", s.panel)
}
