package assert

import (
	"strings"
	"testing"
)

func StringNotContains(
	t *testing.T,
	unexpected string,
	actual string,
) {
	if strings.Contains(actual, unexpected) {
		t.Helper()
		t.Errorf(
			"\nExpect string NOT to contain: %+q\nActual: %+q",
			unexpected,
			actual,
		)
	}
}
