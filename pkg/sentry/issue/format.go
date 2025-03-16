package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (i *Issue) Format(s *format.Settings) string {
	t := status.New(s).String(
		i.Project,
		i.Title,
		i.formatType(s),
		i.formatAge(),
	)
	t.TagLine(tag.Link, "  %s", i.Link)

	return t.Format()
}
