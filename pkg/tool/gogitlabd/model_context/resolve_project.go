package model_context

import (
	"fmt"
	"strconv"
)

func (s *Server) resolveProject(identifier string) (int64, error) {
	id, e := strconv.ParseInt(identifier, 10, 64)

	if e == nil {
		return id, nil
	}

	project, _, e := s.client.Projects.GetProject(identifier, nil)

	if e != nil {
		return 0, fmt.Errorf("project %s: %w", identifier, e)
	}

	return int64(project.ID), nil
}
