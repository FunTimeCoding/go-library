package assert

import "testing"

func Less(
	t *testing.T,
	than float64,
	actual float64,
) {
	if actual >= than {
		t.Helper()
		t.Errorf("\nExpect less than: %f\nActual: %f", than, actual)
	}
}
