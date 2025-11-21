package repository

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Repository) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", r.Name)
	}

	return r.Name
}
