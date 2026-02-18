package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) Start() {
	s.Setup(s.Mux)

	var handler http.Handler = s.Mux

	if s.Middleware != nil {
		handler = s.Middleware(handler)
	}

	s.http = web.Server(handler, s.Address)
	web.ServeAsynchronous(s.http)
}
