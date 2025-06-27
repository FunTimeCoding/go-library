package repository

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (r *Repository) Format(f *option.Format) string {
	return status.New(f).String(
		r.Owner,
		r.Name,
		r.CreatedAt.Format(time.DateMinute),
	).RawList(r).Format()
}
