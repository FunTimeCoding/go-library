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
	if !strings.HasSuffix(actual, expect) {
		t.Helper()
		t.Errorf("expect suffix %q, got %q", expect, actual)
	}
}
