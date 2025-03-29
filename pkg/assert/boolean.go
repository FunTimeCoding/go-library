package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Boolean(
	t *testing.T,
	expect bool,
	actual bool,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
