package run

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (r *Run) Format(f *option.Format) string {
	s := status.New(f).String(
		r.formatName(f),
		r.Status,
	)

	if f.HasTag(tag.Timestamp) {
		s.String(r.Create.Format(time.DateMinute))
	}
	
	s.String(r.formatConcern(f)).RawList(r.Raw)

	return s.Format()
}
