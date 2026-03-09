package image

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"strings"
)

func (i *Image) formatTags(f *option.Format) string {
	if len(i.Tags) == 0 {
		if f.UseColor {
			return console.Yellow("%s", NoTags)
		}

		return NoTags
	}

	return strings.Join(i.Tags, ", ")
}
