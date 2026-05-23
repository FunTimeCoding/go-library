package web

import "fmt"

func (s *Server) usageSummary() []string {
	result := s.service.Usage()

	if result == nil {
		return nil
	}

	return []string{
		fmt.Sprintf("Session %d%%", result.SessionPercent),
		fmt.Sprintf("resets %s", result.SessionReset),
	}
}
