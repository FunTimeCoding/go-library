package errors

import (
	"errors"
	"fmt"
)

func From(a any) error {
	switch cast := a.(type) {
	case error:
		return cast
	default:
		v := fmt.Sprintf("%v", a)

		if v != "" {
			return errors.New(fmt.Sprintf("non-error: %s", v))
		}
	}

	return nil
}
