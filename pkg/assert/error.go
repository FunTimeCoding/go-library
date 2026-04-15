package assert

import "testing"

func Error(
	t *testing.T,
	actual error,
) {
	if actual == nil {
		t.Helper()
		t.Errorf("expected error, got nil")
	}
}
