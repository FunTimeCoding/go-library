package report

import "fmt"

func (s *Section) String(
	name string,
	v string,
) {
	s.appendRenderable(
		&line{
			value:  fmt.Sprintf("%s: %s", name, v),
			indent: s.indent + 1,
		},
	)
}
