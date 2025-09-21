package assert

import (
	"strings"
	"testing"
)

func Prefix(
	t *testing.T,
	expect string,
	actual string,
) {
	t.Helper()

	if !strings.HasPrefix(actual, expect) {
		t.Fatalf("expect prefix %q, got %q", expect, actual)
	}
}
