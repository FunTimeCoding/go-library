package cluster

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Cluster) formatName(f *option.Format) string {
	result := c.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
