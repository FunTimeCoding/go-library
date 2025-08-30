package custom_field_choice

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Choice) formatName(f *option.Format) string {
	result := c.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
