package errors

import "errors"

func Is(e error, target error) bool {
	return errors.Is(e, target)
}
