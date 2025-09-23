package content

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Content) formatTitle(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", c.Title)
	}

	return c.Title
}
