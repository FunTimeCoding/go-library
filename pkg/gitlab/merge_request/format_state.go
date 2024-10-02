package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (r *Request) formatState(s *format.Settings) string {
	if s.UseColor && r.State == OpenedState {
		return console.Yellow(r.State)
	}

	return r.State
}
