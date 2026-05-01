package model_context

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	generative.New(s.server).Setup(m)
}
