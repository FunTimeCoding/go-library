package web

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "User-Agent", UserAgentHeader)
}
