package testutil

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"strings"
	"testing"
)

func AssertBlockedContains(
	t *testing.T,
	results *output.Results,
	substring string,
) {
	t.Helper()

	for _, e := range results.Entries {
		if e.Blocked && strings.Contains(e.Message, substring) {
			return
		}
	}

	t.Errorf("no blocked result containing %q", substring)

	for _, e := range results.Entries {
		if e.Blocked {
			t.Logf("  blocked: %s: %s", e.Path, e.Message)
		}
	}
}
