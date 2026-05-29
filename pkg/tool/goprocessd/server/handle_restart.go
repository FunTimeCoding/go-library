package server

import "fmt"

func (s *Server) handleRestart(arguments []string) string {
	for _, name := range arguments {
		p := s.findProcess(name)

		if p == nil {
			return fmt.Sprintf("error: unknown process %s", name)
		}

		if e := p.Stop(); e != nil {
			return fmt.Sprintf("error: %s", e)
		}

		p.Spawn(s.environment.Build(), nil)
	}

	return "ok"
}
