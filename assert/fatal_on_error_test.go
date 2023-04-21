package assert

import "testing"

func TestFatalOnError(t *testing.T) {
	FatalOnError(t, nil)
}
