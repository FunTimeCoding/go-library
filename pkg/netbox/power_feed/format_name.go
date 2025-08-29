package power_feed

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (f *Feed) formatName(o *option.Format) string {
	result := f.Name

	if o.UseColor {
		result = console.Cyan(result)
	}

	return result
}
