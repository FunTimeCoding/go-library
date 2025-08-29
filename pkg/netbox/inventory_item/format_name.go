package inventory_item

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Item) formatName(f *option.Format) string {
	result := i.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
