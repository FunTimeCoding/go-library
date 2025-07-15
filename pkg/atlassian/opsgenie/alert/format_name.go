package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatName(f *option.Format) string {
	result := a.Name

	if f.UseColor {
		// Not meant to be cyan, because entity and category are cyan
		result = console.Yellow("%s", result)
	}

	return result
}
