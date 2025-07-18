package project

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (p *Project) Format(f *option.Format) string {
	s := status.New(f).Integer(
		p.Identifier,
	).String(p.formatName(f)).RawList(p.Raw)

	if !f.ShowExtended {
		s.TagLine(tag.Link, "  %s", p.Link)
	}

	return s.Format()
}
