package content

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (c *Content) Format(f *option.Format) string {
	s := status.New(f).String(c.formatTitle(f), c.formatSpace())

	if v := c.formatLabels(); v != "" {
		s.String(v)
	}

	if f.HasTag(tag.Dense) {
		if c.TinyLink != "" {
			s.TagLine(tag.Link, "  %s", c.TinyLink)
		}
	} else if c.Link != "" {
		s.TagLine(tag.Link, "  %s", c.Link)
	}

	return s.Format()
}
