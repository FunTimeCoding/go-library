package runner

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (r *Runner) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer(r.Identifier)
	}

	s.String(
		r.formatName(f),
		r.formatDescription(),
		r.hostname(),
	)

	if f.HasTag(tag.Type) {
		s.String(r.Type)
	}

	if f.ShowExtended {
		s.String(r.formatShared())
	}

	s.String(
		r.formatStatus(f),
		r.formatConcern(f),
	).RawList(r.RawList).RawDetail(r.RawDetail)

	if v := r.formatTags(); v != "" {
		s.String(v)
	}

	return s.Format()
}
