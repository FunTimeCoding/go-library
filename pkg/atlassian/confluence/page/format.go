package page

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (p *Page) Format(f *option.Format) string {
	s := status.New(f).String(p.Name)

	if f.HasTag(tag.Dense) {
		s.DetailLink(p.TinyLink, "Confluence", "")
	} else {
		s.DetailLink(p.Link, "Confluence", "")
	}

	return s.Format()
}
