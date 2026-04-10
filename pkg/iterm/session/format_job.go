package session

import "github.com/funtimecoding/go-library/pkg/console/status/option"

func (s *Session) formatJob(_ *option.Format) string {
	if s.JobName == "" {
		return NoJob
	}

	return s.JobName
}
