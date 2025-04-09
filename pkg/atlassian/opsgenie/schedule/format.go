package schedule

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Schedule) Format(t *option.Format) string {
	result := status.New(t).String(s.Name, s.Team.Name).RawList(s.Raw)

	if s.Description != "" {
		result.String(s.Description)
	}

	return result.Format()
}
