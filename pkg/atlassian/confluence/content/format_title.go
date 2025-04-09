package content

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Content) formatTitle(f *option.Format) string {
	result := c.Title

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
