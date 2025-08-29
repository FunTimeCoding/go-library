package power_outlet

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (o *Outlet) formatName(f *option.Format) string {
	result := o.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
