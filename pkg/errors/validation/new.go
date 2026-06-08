package validation

import "fmt"

func New(
	format string,
	a ...any,
) *Detail {
	return &Detail{Message: fmt.Sprintf(format, a...)}
}
