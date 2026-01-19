package assert

import "testing"

func TestStringContains(t *testing.T) {
	StringContains(t, "friend", "hello friend")
}
