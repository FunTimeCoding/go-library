package branch

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (b *Branch) Format(f *format.Settings) string {
	r := status.New(f)
	r.String(b.Name, b.formatMerged(f), b.CommitDate.Format(time.DateMinute))

	return r.Format()
}
