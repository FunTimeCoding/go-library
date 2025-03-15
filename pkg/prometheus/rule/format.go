package rule

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (r *Rule) Format(s *format.Settings) string {
	t := status.New(s).String(
		r.Name,
		r.Group,
		r.formatType(),
	).Raw(r)

	return t.Format()
}
