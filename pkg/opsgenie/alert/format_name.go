package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatName(f *option.Format) string {
	var result string

	if f.UseColor {
		result = console.Cyan(a.Name)
	} else {
		result = a.Name
	}

	return result
}
