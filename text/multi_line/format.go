package multi_line

import "strings"

func (l *MultiLine) Format() string {
	return strings.Join(l.lines, "\n")
}
