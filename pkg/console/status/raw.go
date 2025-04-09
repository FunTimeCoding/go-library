package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
)

func (s *Status) Raw(
	a any,
	title string,
) *Status {
	if !s.format.ShowRaw {
		return s
	}

	r := fmt.Sprintf("%+v", a)

	if s.format.UseColor {
		r = console.Magenta(r)
	}

	return s.Line("  %s: %s", title, r)
}
