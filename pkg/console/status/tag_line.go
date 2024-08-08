package status

import "fmt"

func (s *Status) TagLine(
	tag string,
	format string,
	a ...any,
) *Status {
	s.linesByTag[tag] = append(s.linesByTag[tag], fmt.Sprintf(format, a...))

	return s
}
