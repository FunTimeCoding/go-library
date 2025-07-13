package project

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Project) formatPath(f *option.Format) string {
	result := p.Path

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
