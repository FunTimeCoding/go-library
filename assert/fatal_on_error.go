package assert

import "testing"

func FatalOnError(
	t *testing.T,
	e error,
) {
	if e != nil {
		t.Fatal(e)
	}
}
