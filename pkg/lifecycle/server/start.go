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

	if s.ReadTimeout > 0 {
		s.http.ReadTimeout = s.ReadTimeout
	} else if s.ReadTimeout < 0 {
		s.http.ReadTimeout = 0
	}

	if s.WriteTimeout > 0 {
		s.http.WriteTimeout = s.WriteTimeout
	} else if s.WriteTimeout < 0 {
		s.http.WriteTimeout = 0
	}

	web.ServeAsynchronous(s.http)
}
