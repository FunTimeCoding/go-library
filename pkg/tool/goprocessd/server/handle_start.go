package server

import "fmt"

func (s *Server) handleStart(arguments []string) string {
	for _, name := range arguments {
		p := s.findProcess(name)

		if p == nil {
			return fmt.Sprintf("error: unknown process %s", name)
		}

		p.Spawn(s.environment.Build(), nil)
	}

	return "ok"
}
