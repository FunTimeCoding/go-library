package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(
		i.FormatStatus(),
		i.FormatSummary(f),
		i.FormatDescription(),
	)
	s.TagLine(tag.Link, "  %s", i.Link)
	s.RawList(i)

	return s.Format()
}
