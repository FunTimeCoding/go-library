package errors

import (
	"context"
	"errors"
)

func Deadline(e error) bool {
	return errors.Is(e, context.DeadlineExceeded)
}
