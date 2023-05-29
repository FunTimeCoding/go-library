package assert

import "testing"

func False(
	t *testing.T,
	actual bool,
) {
	t.Helper()
	Boolean(t, false, actual)
}
