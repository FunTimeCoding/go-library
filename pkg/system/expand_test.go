package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.True(t, len(Expand(constant.Tilde)) > 0)
}
