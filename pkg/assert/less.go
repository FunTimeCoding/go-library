package assert

import "testing"

func Less(
	t *testing.T,
	than float64,
	actual float64,
) {
	t.Helper()

	if actual >= than {
		t.Logf("\nExpect less than: %f\nActual: %f", than, actual)
		t.Fail()
	}
}
