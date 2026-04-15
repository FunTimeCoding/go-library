package assert

import "testing"

func TestStringNotContains(t *testing.T) {
	StringNotContains(t, "enemy", "hello friend")
}
