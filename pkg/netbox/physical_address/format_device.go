package physical_address

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Address) formatDevice(f *option.Format) string {
	if a.Interface == nil {
		return ""
	}

	result := a.Interface.Device.GetName()

	if result == "" {
		result = NoDevice

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
