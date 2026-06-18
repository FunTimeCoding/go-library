package server

import "net/http"

func (s *Server) WithMiddleware(
	middleware func(http.Handler) http.Handler,
) *Server {
	s.Middleware = middleware

	return s
}
