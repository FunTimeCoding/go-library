package validation

import "fmt"

func New(
	format string,
	a ...any,
) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
