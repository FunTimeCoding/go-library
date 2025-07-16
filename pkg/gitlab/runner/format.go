package runner

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Runner) Format(f *option.Format) string {
	s := status.New(f).Integer(r.Identifier).String(
		r.formatName(f),
		r.formatDescription(),
		r.hostname(),
		r.Type,
		r.formatShared(),
		r.formatStatus(f),
		r.formatConcern(f),
	).RawList(r.RawList).RawDetail(r.RawDetail)

	if v := r.formatTags(); v != "" {
		s.String(v)
	}

	return s.Format()
}
