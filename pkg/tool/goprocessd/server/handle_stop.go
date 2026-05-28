package server

import "fmt"

func (s *Server) handleStop(arguments []string) string {
	for _, name := range arguments {
		p := s.findProcess(name)

		if p == nil {
			return fmt.Sprintf("error: unknown process %s", name)
		}

		if e := p.Stop(); e != nil {
			return fmt.Sprintf("error: %s", e)
		}
	}

	return "ok"
}
