package custom_link

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (l *Link) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", l.Name)
	}

	return l.Name
}
