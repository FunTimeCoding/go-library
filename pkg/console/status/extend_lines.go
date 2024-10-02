package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func extendLines(
	input string,
	lines []string,
	indent int,
) string {
	return fmt.Sprintf(
		"%s%s\n%s",
		spaces(indent),
		input,
		join.NewLine(lines),
	)
}
