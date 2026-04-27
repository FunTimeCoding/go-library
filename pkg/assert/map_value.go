package assert

import (
	"fmt"
	"testing"
)

func MapValue(
	t *testing.T,
	expect any,
	m map[string]any,
	key string,
) {
	t.Helper()
	actual, okay := m[key]

	if !okay {
		t.Errorf("key %q: not found in map", key)

		return
	}

	expectString := fmt.Sprintf("%v", expect)
	actualString := fmt.Sprintf("%v", actual)

	if expectString != actualString {
		t.Errorf(
			"key %q: expected %v (%T), got %v (%T)",
			key,
			expect,
			expect,
			actual,
			actual,
		)
	}
}
