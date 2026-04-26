package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	assert.True(t, len(AbsolutePath(constant.CurrentDirectory)) > 0)
}
