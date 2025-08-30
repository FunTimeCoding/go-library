package virtual_network

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (n *Network) formatName(f *option.Format) string {
	result := n.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
