package device

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (d *Device) formatName(f *option.Format) string {
	result := d.Name

	if result == "" {
		result = NoName

		if f.UseColor {
			result = console.Yellow(result)
		}

		return result
	}

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
