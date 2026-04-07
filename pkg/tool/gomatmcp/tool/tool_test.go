package tool

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestTool(t *testing.T) {
	assert.NotNil(t, New(nil, nil))
}
