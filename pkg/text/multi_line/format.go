package multi_line

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func (l *MultiLine) Format() string {
	return strings.Join(l.lines, separator.Unix)
}