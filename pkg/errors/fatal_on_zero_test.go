package errors

import "testing"

func TestFatalOnZero(t *testing.T) {
	FatalOnZero(1, "Example")
}
