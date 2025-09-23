package module_bay

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Bay) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", b.Name)
	}

	return b.Name
}
