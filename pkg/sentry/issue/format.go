package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(
		i.Project,
		i.Title,
		i.formatType(f),
		i.formatAge(),
	)
	s.TagLine(tag.Link, "  %s", i.Link)

	return s.Format()
}
