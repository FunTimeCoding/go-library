package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(
		i.FormatStatus(),
		i.FormatSummary(f),
	)

	if d := i.FormatDescription(f); d != NoDescription {
		s.Line("%s", d)
	}

	s.DetailLink(i.Link, "Jira", "")
	s.RawList(i)

	return s.Format()
}
