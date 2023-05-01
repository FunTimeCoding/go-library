package report

import (
	"fmt"
	"github.com/funtimecoding/go-library/floats"
)

func (s *Section) Percent(
	name string,
	f float64,
) {
	s.appendRenderable(
		&line{
			value: fmt.Sprintf(
				"%s: %s%%",
				name,
				floats.ToStringRounded(f),
			),
			indent: s.indent + 1,
		},
	)
}
