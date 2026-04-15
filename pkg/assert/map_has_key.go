package assert

import "testing"

func MapHasKey(
	t *testing.T,
	m map[string]any,
	key string,
) {
	t.Helper()

	if _, ok := m[key]; !ok {
		t.Errorf("expected key %q in map, not found", key)
	}
}
