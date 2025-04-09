package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Request) formatState(f *option.Format) string {
	result := r.State

	if result == OpenedState {
		result = OpenAlias
	}

	if f.UseColor {
		if result == OpenAlias {
			return console.Yellow("%s", result)
		}

		if result == ClosedState {
			return console.Green("%s", result)
		}
	}

	return result
}
