package assert

import "testing"

func TestCount(t *testing.T) {
	Count(t, 0, []string{})
	Count(t, 1, []string{"1"})
	Count(t, 2, []string{"1", "2"})
}
