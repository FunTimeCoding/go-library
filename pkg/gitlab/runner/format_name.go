package runner

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Runner) formatName(f *option.Format) string {
	result := r.Name

	if result == "" {
		result = NoName
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
