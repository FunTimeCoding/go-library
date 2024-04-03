package errors

import "testing"

func TestUnexpectedCount(t *testing.T) {
	UnexpectedCount(1, 1)
}
