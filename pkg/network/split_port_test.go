package network

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestSplitPort(t *testing.T) {
	assert.Equal(t, 80, SplitPort("127.0.0.1:80"))
}
