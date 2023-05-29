package assert

import "testing"

func True(
	t *testing.T,
	actual bool,
) {
	t.Helper()
	Boolean(t, true, actual)
}
