package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func String(
	t *testing.T,
	expected string,
	actual string,
) {
	t.Helper()
	assert.Equal(t, actual, expected)
}
