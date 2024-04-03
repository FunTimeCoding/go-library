package assert

import (
	"bytes"
	"gotest.tools/v3/assert"
	"testing"
)

func Bytes(
	t *testing.T,
	expected []byte,
	actual []byte,
) {
	t.Helper()
	assert.Assert(
		t,
		bytes.Equal(actual, expected),
		"expected %v, got %v",
		expected,
		actual,
	)
}
