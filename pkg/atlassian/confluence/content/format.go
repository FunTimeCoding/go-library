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
		s.DetailLink(c.TinyLink, "Confluence", "")
	} else {
		s.DetailLink(c.Link, "Confluence", "")
	}

	return s.Format()
}
