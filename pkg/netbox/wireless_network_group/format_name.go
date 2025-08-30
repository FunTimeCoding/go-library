package wireless_network_group

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (g *Group) formatName(f *option.Format) string {
	result := g.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
