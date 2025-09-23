package cable

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Cable) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", c.Name)
	}

	return c.Name
}
