package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatEntity(f *option.Format) string {
	result := a.Entity

	if result != "" && f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
