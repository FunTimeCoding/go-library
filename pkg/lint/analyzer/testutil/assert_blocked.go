package testutil

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"testing"
)

func AssertBlocked(
	t *testing.T,
	results *output.Results,
	count int,
) {
	t.Helper()
	blocked := countBlocked(results)

	if blocked != count {
		t.Errorf("expected %d blocked, got %d", count, blocked)

		for _, e := range results.Entries {
			if e.Blocked {
				t.Logf("  blocked: %s: %s", e.Path, e.Message)
			}
		}
	}
}
