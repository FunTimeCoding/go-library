package conversations

import "net/http"

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /conversations", s.page)
	m.HandleFunc("GET /conversations/sidebar", s.sidebar)
	m.HandleFunc("GET /conversations/{identifier}/edit", s.editForm)
	m.HandleFunc("POST /conversations/{identifier}/edit", s.editSubmit)
	m.HandleFunc("GET /conversations/{identifier}", s.panel)
}
