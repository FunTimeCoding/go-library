package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (a *Alert) formatEntity(s *format.Settings) string {
	result := a.Entity

	if result != "" && s.UseColor {
		result = console.Cyan(result)
	}

	return result
}
