package system_number

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (n *Number) formatName(f *option.Format) string {
	result := n.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
