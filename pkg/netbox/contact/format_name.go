package contact

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Contact) formatName(f *option.Format) string {
	result := c.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
