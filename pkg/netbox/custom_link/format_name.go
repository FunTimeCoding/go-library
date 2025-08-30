package custom_link

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (l *Link) formatName(f *option.Format) string {
	result := l.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
