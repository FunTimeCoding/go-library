package assert

import "testing"

func TestPrefix(t *testing.T) {
	Prefix(t, "", "ab")
	Prefix(t, "a", "ab")
	Prefix(t, "ab", "ab")
}
