package assert

import "testing"

func FatalOnError(
	t *testing.T,
	fail error,
) {
	if fail != nil {
		t.Fatal(fail)
	}
}
