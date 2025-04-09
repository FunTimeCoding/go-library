package run

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (r *Run) Format(f *option.Format) string {
	return status.New(f).String(
		r.Name,
		r.Status,
		r.CreatedAt.Format(time.DateMinute),
	).RawList(r).Format()
}
