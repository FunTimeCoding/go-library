package errors

import "context"

func Deadline(e error) bool {
	return Is(e, context.DeadlineExceeded)
}
