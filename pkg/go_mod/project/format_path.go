package project

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Project) formatPath(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", p.Path)
	}

	return p.Path
}
