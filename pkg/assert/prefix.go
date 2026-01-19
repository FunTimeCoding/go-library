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
	if !strings.HasPrefix(actual, expect) {
		t.Helper()
		t.Errorf("expect prefix %q, got %q", expect, actual)
	}
}
