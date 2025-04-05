package system

import "testing"

func TestExitOnEmpty(t *testing.T) {
	ExitOnEmpty(1, "notEmpty", "This should not exit")
}
