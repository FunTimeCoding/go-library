package project

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Project) formatName(f *option.Format) string {
	result := p.CombinedName()

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
