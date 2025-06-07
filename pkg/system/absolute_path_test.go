package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	assert.True(t, len(AbsolutePath(".")) > 0)
}
