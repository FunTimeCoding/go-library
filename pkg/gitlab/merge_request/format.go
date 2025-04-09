package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (r *Request) Format(f *option.Format) string {
	s := status.New(f).Integer(
		r.Project,
		r.Identifier,
	).String(
		r.formatState(f),
		r.formatTitle(f),
		r.formatAge(f),
	).RawList(r.Raw)

	if !f.ShowExtended {
		s.TagLine(tag.Link, "  %s", r.Link)
	}

	return s.Format()
}
