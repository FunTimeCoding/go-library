package session

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/iterm/constant"
)

func (s *Session) formatJob(_ *option.Format) string {
	if s.JobName == "" {
		return constant.NoJob
	}

	return s.JobName
}
