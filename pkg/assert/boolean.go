package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Boolean(
	t *testing.T,
	expected bool,
	actual bool,
) {
	t.Helper()
	assert.Equal(t, actual, expected)
}
