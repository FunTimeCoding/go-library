package indent

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func F(
	format string,
	indent int,
	a ...any,
) {
	fmt.Printf(
		"%s%s\n",
		strings.Repeat(separator.DoubleSpace, indent),
		fmt.Sprintf(format, a...),
	)
}
