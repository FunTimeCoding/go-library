package image

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Image) formatDigest(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Digest)
	}

	return i.Digest
}
