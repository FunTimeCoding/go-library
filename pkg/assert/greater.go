package assert

import "testing"

func Greater(
	t *testing.T,
	than float64,
	actual float64,
) {
	if actual <= than {
		t.Helper()
		t.Errorf("\nExpect greater than: %f\nActual: %f", than, actual)
	}
}
