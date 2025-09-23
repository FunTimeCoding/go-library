package issue

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Issue) FormatSummary(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Summary)
	}

	return i.Summary
}
