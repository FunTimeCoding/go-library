package repository

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Repository) Format(f *option.Format) string {
	return status.New(f).String(
		r.Path,
		r.formatConcern(f),
	).RawList(r).Format()
}
