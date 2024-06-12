package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Integer64(
	t *testing.T,
	expected int64,
	actual int64,
) {
	t.Helper()
	assert.Equal(t, actual, expected)
}
