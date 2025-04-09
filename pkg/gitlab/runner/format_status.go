package runner

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Runner) formatStatus(f *option.Format) string {
	result := r.Status

	if f.UseColor {
		if result == OnlineStatus {
			result = console.Green("%s", result)
		} else {
			result = console.Red("%s", result)
		}
	}

	return result
}
