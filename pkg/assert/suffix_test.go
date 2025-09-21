package assert

import "testing"

func TestSuffix(t *testing.T) {
	Suffix(t, "", "ab")
	Suffix(t, "b", "ab")
	Suffix(t, "ab", "ab")
}
