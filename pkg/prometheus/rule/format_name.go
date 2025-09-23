package rule

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Rule) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", r.Name)
	}

	return r.Name
}
