package multi_line

import "github.com/funtimecoding/go-library/pkg/strings"

func (l *MultiLine) Format() string {
	return strings.JoinNewLine(l.lines)
}
