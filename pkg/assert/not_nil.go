package assert

import "testing"

func NotNil(
	t *testing.T,
	actual any,
) {
	if actual == nil {
		t.Helper()
		t.Errorf("expected not nil, got nil")
	}
}
