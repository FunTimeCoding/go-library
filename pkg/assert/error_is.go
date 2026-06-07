package assert

import (
	"errors"
	"testing"
)

func ErrorIs(
	t *testing.T,
	actual error,
	expected error,
) {
	t.Helper()

	if !errors.Is(actual, expected) {
		t.Errorf("expected error %v, got %v", expected, actual)
	}
}
