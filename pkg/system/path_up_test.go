package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"runtime"
	"testing"
)

func TestPathUp(t *testing.T) {
	switch runtime.GOOS {
	case constant.Linux:
		assert.String(t, "/a", PathUp("/a/b", 1))
	case constant.Windows:
		assert.String(t, "\\a", PathUp("/a/b", 1))
	}
}
