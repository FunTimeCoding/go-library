package device

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (d *Device) formatName(f *option.Format) string {
	if d.Name == "" {
		if f.UseColor {
			return console.Yellow("%s", NoName)
		}

		return NoName
	}

	if f.UseColor {
		return console.Cyan("%s", d.Name)
	}

	return d.Name
}
