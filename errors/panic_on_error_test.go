package errors

import "testing"

func TestPanicOnError(t *testing.T) {
	PanicOnError(nil)
}
