package multi_line

import "github.com/funtimecoding/go-library/pkg/strings/join"

func (m *MultiLine) Format() string {
	return join.NewLine(m.lines)
}
