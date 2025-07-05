package repository

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (r *Repository) formatConcern(o *option.Format) string {
	if len(r.concern) == 0 {
		return ""
	}

	result := join.Comma(r.concern)

	if o.UseColor {
		console.Yellow("%s", result)
	}

	return result
}
