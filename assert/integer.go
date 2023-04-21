package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Integer(
	t *testing.T,
	expected int,
	actual int,
) {
	t.Helper()
	assert.Equal(t, actual, expected)
}
