package schedule

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Schedule) Format(t *option.Format) string {
	r := status.New(t).String(s.Name, s.Team.Name).RawList(s.Raw)

	if s.Description != "" {
		r.String(s.Description)
	}

	return r.Format()
}
