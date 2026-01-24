package server

import "github.com/funtimecoding/go-library/pkg/web"

func (s *Server) Stop() {
	if s.web != nil {
		web.GracefulShutdown(s.context, s.web, true)
		s.web = nil
	}
}
