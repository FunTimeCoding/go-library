package report

import "fmt"

func (s *Section) Integer(
	name string,
	i int,
) {
	s.appendRenderable(
		&line{
			value:  fmt.Sprintf("%s: %d", name, i),
			indent: s.indent + 1,
		},
	)
}
