package server

func (s *Server) WithProtected() *Server {
	s.protected = true

	return s
}
