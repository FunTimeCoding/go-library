package assert

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Integer32(
	t *testing.T,
	expect int32,
	actual int32,
) {
	t.Helper()
	assert.Equal(t, actual, expect)
}
