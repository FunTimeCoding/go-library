package location

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (l *Location) formatName(f *option.Format) string {
	result := l.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
