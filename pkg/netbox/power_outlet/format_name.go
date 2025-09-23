package power_outlet

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (o *Outlet) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", o.Name)
	}

	return o.Name
}
