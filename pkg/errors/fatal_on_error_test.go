package errors

import "testing"

func TestFatalOnError(t *testing.T) {
	FatalOnError(nil)
}
