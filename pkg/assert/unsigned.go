package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Unsigned(
	t *testing.T,
	expect uint,
	actual uint,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
