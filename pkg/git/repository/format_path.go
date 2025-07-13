package repository

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Repository) formatPath(f *option.Format) string {
	result := r.Path

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
