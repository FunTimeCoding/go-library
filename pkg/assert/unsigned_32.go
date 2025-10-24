package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Unsigned32(
	t *testing.T,
	expect uint32,
	actual uint32,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
