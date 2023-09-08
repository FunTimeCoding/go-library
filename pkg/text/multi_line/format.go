package multi_line

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (l *MultiLine) Format() string {
	return join.NewLine(l.lines)
}
