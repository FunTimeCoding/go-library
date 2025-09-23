package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	if a.Category == "" {
		return ""
	}

	if f.UseColor {
		return console.Cyan("%s", a.Category)
	}

	return a.Category
}
