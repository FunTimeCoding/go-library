package repository

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Repository) formatPath(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", r.Path)
	}

	return r.Path
}
