package inventory_item

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Item) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Name)
	}

	return i.Name
}
