package page

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (p *Page) Format(f *option.Format) string {
	s := status.New(f).String(p.Name)

	if f.HasTag(tag.Dense) {
		if p.TinyLink != "" {
			s.TagLine(tag.Link, "  %s", p.TinyLink)
		}
	} else if p.Link != "" {
		s.TagLine(tag.Link, "  %s", p.Link)
	}

	return s.Format()
}
