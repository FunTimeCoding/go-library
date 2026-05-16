package model_context

import (
	"fmt"
	"strconv"
)

func (s *Server) resolveProject(identifier string) (int64, error) {
	i, e := strconv.ParseInt(identifier, 10, 64)

	if e == nil {
		return i, nil
	}

	project, _, f := s.client.Projects.GetProject(identifier, nil)

	if f != nil {
		return 0, fmt.Errorf("project %s: %w", identifier, f)
	}

	return project.ID, nil
}
