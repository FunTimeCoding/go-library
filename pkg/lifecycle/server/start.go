package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) Start() {
	s.Setup(s.Mux)
	var m http.Handler = s.Mux

	if s.Middleware != nil {
		m = s.Middleware(m)
	}

	s.http = web.Server(m, s.Address)
	s.http.ReadTimeout = s.ReadTimeout
	s.http.WriteTimeout = s.WriteTimeout
	web.ServeAsynchronous(s.http)
}
