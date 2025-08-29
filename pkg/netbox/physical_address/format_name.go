package physical_address

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Address) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", a.Name)
	}

	return a.Name
}
