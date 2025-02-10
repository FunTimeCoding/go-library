package assert

import "testing"

func FatalOnError(
	t *testing.T,
	e error,
) {
	t.Helper()

	if e != nil {
		t.Fatal(e)
	}
}
