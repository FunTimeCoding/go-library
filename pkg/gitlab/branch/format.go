package branch

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Branch) Format(f *option.Format) string {
	s := status.New(f)
	s.String(b.Name, b.formatMerged(f), b.formatAge(f))

	return s.Format()
}
