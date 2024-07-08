package status

import "fmt"

func (s *Status) Line(
	format string,
	a ...any,
) *Status {
	s.lines = append(s.lines, fmt.Sprintf(format, a...))

	return s
}
