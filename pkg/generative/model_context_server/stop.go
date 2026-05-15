package model_context_server

func (s *Server) Stop() {
	s.Lifecycle.Stop()
}
