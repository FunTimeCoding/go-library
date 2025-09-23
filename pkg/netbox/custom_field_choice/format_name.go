package custom_field_choice

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Choice) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", c.Name)
	}

	return c.Name
}
