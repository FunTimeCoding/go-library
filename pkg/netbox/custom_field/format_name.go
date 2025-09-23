package custom_field

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (f *Field) formatName(o *option.Format) string {
	if o.UseColor {
		return console.Cyan("%s", f.Name)
	}

	return f.Name
}
