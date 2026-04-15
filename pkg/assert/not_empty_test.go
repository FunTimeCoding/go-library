package assert

import "testing"

func TestNotEmpty(t *testing.T) {
	NotEmpty(t, []string{"one"})
	NotEmpty(t, map[string]int{"a": 1})
}
