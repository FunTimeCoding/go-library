package project

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (p *Project) Format(s *format.Settings) string {
	f := status.New(s).Integer(
		p.Identifier,
	).String(p.CombinedName()).Raw(p.Raw)

	if !s.ShowExtended {
		f.TagLine(status.LinkTag, "  %s", p.Link)
	}

	return f.Format()
}
