package server

import (
	"fmt"
	"strings"
)

func (s *Server) handleLog(arguments []string) string {
	if len(arguments) == 0 {
		return "error: log requires a process name"
	}

	p := s.findProcess(arguments[0])

	if p == nil {
		return fmt.Sprintf("error: unknown process %s", arguments[0])
	}

	lines := p.Log()

	if len(lines) == 0 {
		return "ok"
	}

	return strings.Join(lines, "\n")
}
