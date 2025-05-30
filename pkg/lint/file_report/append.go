package file_report

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func (r *Report) append(
	b *strings.Builder,
	line string,
) {
	b.WriteString(line)
	b.WriteString(separator.Unix)
}
