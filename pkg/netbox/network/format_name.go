package network

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Interface) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Name)
	}

	return i.Name
}
