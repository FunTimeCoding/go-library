package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (r *Request) Format(s *format.Settings) string {
	f := status.New(s).Integer(
		r.Project,
		r.Identifier,
	).String(r.formatState(s), r.formatTitle(s)).Raw(r.Raw)

	if !s.ShowExtended {
		f.TagLine(status.LinkTag, "  %s", r.Link)
	}

	return f.Format()
}
