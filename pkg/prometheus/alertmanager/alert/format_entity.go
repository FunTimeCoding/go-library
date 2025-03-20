package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (a *Alert) formatCategory(s *format.Settings) string {
	result := a.Category

	if result != "" && s.UseColor {
		result = console.Cyan(result)
	}

	return result
}
