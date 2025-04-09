package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	result := a.Category

	if result != "" && f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
