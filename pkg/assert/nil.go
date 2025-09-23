package assert

import "testing"

func Nil(
	t *testing.T,
	actual any,
) {
	t.Helper()

	if actual != nil {
		t.Errorf("expected nil, got %T", actual)
	}
}
