package contact_group

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (g *Group) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", g.Name)
	}

	return g.Name
}
