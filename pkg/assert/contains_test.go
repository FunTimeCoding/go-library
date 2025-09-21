package assert

import "testing"

func TestContains(t *testing.T) {
	Contains(t, []string{}, []string{"a", "b"})
	Contains(t, []string{"a"}, []string{"a", "b"})
	Contains(t, []string{"a", "b"}, []string{"a", "b"})
}
