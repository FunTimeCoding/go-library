package branch

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (b *Branch) Format(f *format.Settings) string {
	r := status.New(f)
	r.String(b.Name, b.formatMerged(f), b.formatAge())

	return r.Format()
}
