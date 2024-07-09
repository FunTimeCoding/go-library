package status

import "fmt"

func (s *Status) Raw(a any) *Status {
	r := fmt.Sprintf("%+v", a)

	if s.color {
		s.Line("  Raw: %s", magenta(r))
	} else {
		s.Line("  Raw: %s", r)
	}

	return s
}
