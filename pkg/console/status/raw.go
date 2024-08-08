package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
)

func (s *Status) Raw(a any) *Status {
	if !s.settings.ShowRaw {
		return s
	}

	r := fmt.Sprintf("%+v", a)

	if s.settings.UseColor {
		s.Line("  Raw: %s", console.Magenta(r))
	} else {
		s.Line("  Raw: %s", r)
	}

	return s
}
