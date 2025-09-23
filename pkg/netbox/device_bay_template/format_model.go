package device_bay_template

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Template) formatModel(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", t.Model)
	}

	return t.Model
}
