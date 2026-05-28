package server

import "strings"

func (s *Server) handleList() string {
	names := make([]string, len(s.processes))

	for i, p := range s.processes {
		names[i] = p.Name
	}

	return strings.Join(names, "\n")
}
