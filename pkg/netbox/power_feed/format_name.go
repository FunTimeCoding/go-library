package power_feed

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (f *Feed) formatName(o *option.Format) string {
	if o.UseColor {
		return console.Cyan("%s", f.Name)
	}

	return f.Name
}
