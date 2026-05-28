package server

import "fmt"

func (s *Server) handleReloadEnvironment() string {
	if e := s.environment.Load(s.envrcPath); e != nil {
		return fmt.Sprintf("error: %s", e)
	}

	return "ok"
}
