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
		if l := c.TinyLink; l != "" {
			s.TagLine(tag.Link, "  %s", l)
		}
	} else if l := c.Link; l != "" {
		s.TagLine(tag.Link, "  %s", l)
	}

	return s.Format()
}
