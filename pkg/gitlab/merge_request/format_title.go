package merge_request

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (r *Request) formatTitle(s *format.Settings) string {
	if s.UseColor {
		return console.Cyan(r.Title)
	}

	return r.Title
}
