package assert

import (
	"testing"
)

func Greater(
	t *testing.T,
	than float64,
	actual float64,
) {
	t.Helper()

	if actual <= than {
		t.Fail()
	}
}
