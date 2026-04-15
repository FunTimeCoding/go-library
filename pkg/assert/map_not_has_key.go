package assert

import "testing"

func MapNotHasKey(
	t *testing.T,
	m map[string]any,
	key string,
) {
	t.Helper()

	if _, ok := m[key]; ok {
		t.Errorf("expected key %q NOT in map, but found", key)
	}
}
