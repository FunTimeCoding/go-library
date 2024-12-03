package project

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (p *Project) Format(s *format.Settings) string {
	f := status.New(s).Integer(
		p.Identifier,
	).String(p.CombinedName()).Raw(p.Raw)

	if !s.ShowExtended {
		f.TagLine(tag.Link, "  %s", p.Link)
	}

	return f.Format()
}
