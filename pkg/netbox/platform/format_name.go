package platform

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Platform) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", p.Name)
	}

	return p.Name
}
