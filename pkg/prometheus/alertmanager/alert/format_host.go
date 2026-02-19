package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatHost(f *option.Format) string {
	result := a.Host()

	if result == "" {
		result = NoHost

		if f.UseColor {
			result = console.Yellow(result)
		}

		return result
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
