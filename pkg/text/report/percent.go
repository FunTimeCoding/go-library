package report

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/floats"
	"github.com/funtimecoding/go-library/pkg/text/report/line"
)

func (s *Section) Percent(
	name string,
	f float64,
) {
	l := line.New()
	l.Value = fmt.Sprintf(
		"%s: %s%%",
		name,
		floats.ToStringRounded(f),
	)
	l.Indent = s.indent + 1
	s.appendRenderable(l)
}
