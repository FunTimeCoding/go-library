package web

import (
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.heatmap)
	m.HandleFunc("GET /events", s.events)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
