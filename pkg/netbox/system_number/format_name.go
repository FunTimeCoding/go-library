package system_number

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (n *Number) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", n.Name)
	}

	return n.Name
}
