package assert

import "testing"

func TestStrings(t *testing.T) {
	Strings(t, []string{"a", "b"}, []string{"a", "b"})
}
