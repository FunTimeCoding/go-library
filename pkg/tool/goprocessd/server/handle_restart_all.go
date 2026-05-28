package server

import "fmt"

func (s *Server) handleRestartAll() string {
	for _, p := range s.processes {
		if e := p.Stop(); e != nil {
			return fmt.Sprintf("error: %s", e)
		}

		p.Spawn(s.environment.Build(), nil, nil)
	}

	return "ok"
}
