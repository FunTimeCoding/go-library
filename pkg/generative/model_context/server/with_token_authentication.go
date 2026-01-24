package server

func WithTokenAuthentication(token string) Option {
	return func(s *Server) {
		s.tokenAuthentication = true
		s.token = token
	}
}
