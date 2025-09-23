package tenant

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Tenant) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", t.Name)
	}

	return t.Name
}
