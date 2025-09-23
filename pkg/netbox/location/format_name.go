package location

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (l *Location) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", l.Name)
	}

	return l.Name
}
