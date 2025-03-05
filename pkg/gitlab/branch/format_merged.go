package branch

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (b *Branch) formatMerged(f *format.Settings) string {
	var result string

	if b.Merged {
		result = Merged

		if f.UseColor {
			result = console.Green(result)
		}
	} else {
		result = Unmerged

		if f.UseColor {
			result = console.Red(result)
		}
	}

	return result
}
