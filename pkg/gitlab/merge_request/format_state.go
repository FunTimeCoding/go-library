package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (r *Request) formatState(s *format.Settings) string {
	result := r.State

	if result == OpenedState {
		result = OpenAlias
	}

	if s.UseColor {
		if result == OpenAlias {
			return console.Yellow(result)
		}

		if result == ClosedState {
			return console.Green(result)
		}
	}

	return result
}
