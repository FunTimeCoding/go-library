package assert

import (
	"bytes"
	"gotest.tools/v3/assert"
	"testing"
)

func Bytes(
	t *testing.T,
	expect []byte,
	actual []byte,
) {
	t.Helper()
	assert.Assert(
		t,
		bytes.Equal(actual, expect),
		"expect %v, got %v",
		expect,
		actual,
	)
}
