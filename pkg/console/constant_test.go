package console

import "testing"

func TestConstant(t *testing.T) {
	Blue("%s", "test")
	Cyan("%s", "test")
	Green("%s", "test")
	Magenta("%s", "test")
	Red("%s", "test")
	Yellow("%s", "test")
}
