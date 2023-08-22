package report

import "fmt"

func (s *Section) Float(
	name string,
	f float64,
	unit string,
) {
	var result string

	if unit == "" {
		result = fmt.Sprintf("%s: %.1f", name, f)
	} else {
		result = fmt.Sprintf("%s: %.1f %s", name, f, unit)
	}

	s.appendRenderable(
		&line{
			value:  result,
			indent: s.indent + 1,
		},
	)
}
