package assert

import "testing"

func LessEqual(
	t *testing.T,
	than float64,
	actual float64,
) {
	t.Helper()

	if actual > than {
		t.Logf("\nExpected less equal than: %f\nActual: %f", than, actual)
		t.Fail()
	}
}
