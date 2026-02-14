package network

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestSplitHostPort(t *testing.T) {
	host, port := SplitHostPort("127.0.0.1:80")
	assert.Equal(t, "127.0.0.1", host)
	assert.Equal(t, 80, port)
}
