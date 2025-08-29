package device_role

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (d *DeviceRole) formatName(f *option.Format) string {
	result := d.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
