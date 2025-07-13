package project

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Project) Format(f *option.Format) string {
	return status.New(f).String(
		p.formatPath(f),
		p.formatConcern(f),
	).RawList(p).Format()
}
