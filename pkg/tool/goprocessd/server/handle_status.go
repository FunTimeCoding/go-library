package server

import (
	"fmt"
	"strings"
)

func (s *Server) handleStatus() string {
	var lines []string

	for _, p := range s.processes {
		prefix := " "

		if p.Running() {
			prefix = "*"
		}

		lines = append(lines, fmt.Sprintf("%s%s", prefix, p.Name))
	}

	return strings.Join(lines, "\n")
}
