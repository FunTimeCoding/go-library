package server

import "github.com/funtimecoding/go-library/pkg/tool/goprocessd/process"

func (s *Server) findProcess(name string) *process.Process {
	for _, p := range s.processes {
		if p.Name == name {
			return p
		}
	}

	return nil
}
