package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.True(t, len(Expand("~")) > 0)
}
