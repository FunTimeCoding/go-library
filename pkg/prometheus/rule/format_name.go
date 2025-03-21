package rule

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (r *Rule) formatName(s *format.Settings) string {
	result := r.Name

	if s.UseColor {
		result = console.Cyan(result)
	}

	return result
}
