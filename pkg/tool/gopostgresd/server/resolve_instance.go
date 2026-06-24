package server

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) resolveInstance(name string) (*server.Error, bool) {
	if _, okay := s.store.Instance(name); okay {
		return nil, true
	}

	return &server.Error{
		Error: fmt.Sprintf("instance not found: %s", name),
	}, false
}
