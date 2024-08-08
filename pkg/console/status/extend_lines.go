package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func extendLines(
	input string,
	lines []string,
) string {
	return fmt.Sprintf("%s\n%s", input, join.NewLine(lines))
}
