package assert

import "testing"

func NotNil(
	t *testing.T,
	actual any,
) {
	t.Helper()

	if actual == nil {
		t.Errorf("expected not nil, got nil")
	}
}
