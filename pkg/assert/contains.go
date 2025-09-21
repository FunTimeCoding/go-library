package assert

import (
	"slices"
	"testing"
)

func Contains(
	t *testing.T,
	expect []string,
	actual []string,
) {
	t.Helper()

	for _, e := range expect {
		if !slices.Contains(actual, e) {
			t.Errorf("\nExpect contains: %+q\nActual: %+q", e, actual)

			return
		}
	}
}
