package branch

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Branch) formatMerged(f *option.Format) string {
	var result string

	if b.Merged {
		result = Merged

		if f.UseColor {
			return console.Green(result)
		}
	} else {
		result = Unmerged

		if f.UseColor {
			return console.Yellow(result)
		}
	}

	return result
}
