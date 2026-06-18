package server

func (s *Server) WithProfiling() *Server {
	s.profiling = true

	return s
}
