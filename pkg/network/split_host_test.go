package network

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestSplitHost(t *testing.T) {
	assert.Equal(t, "127.0.0.1", SplitHost("127.0.0.1:80"))
}
