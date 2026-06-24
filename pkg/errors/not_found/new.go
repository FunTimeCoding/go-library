package not_found

import "fmt"

func New(
	kind string,
	identifier any,
) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf("%s not found: %v", kind, identifier),
	}
}
