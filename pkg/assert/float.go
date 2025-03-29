package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Float(
	t *testing.T,
	expect float64,
	actual float64,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
