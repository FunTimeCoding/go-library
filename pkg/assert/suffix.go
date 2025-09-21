package assert

import (
	"strings"
	"testing"
)

func Suffix(
	t *testing.T,
	expect string,
	actual string,
) {
	t.Helper()

	if !strings.HasSuffix(actual, expect) {
		t.Fatalf("expect suffix %q, got %q", expect, actual)
	}
}
