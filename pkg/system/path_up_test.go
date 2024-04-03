package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"runtime"
	"testing"
)

func TestPathUp(t *testing.T) {
	switch runtime.GOOS {
	case Linux:
		assert.String(t, "/a", PathUp("/a/b", 1))
	case Windows:
		assert.String(t, "\\a", PathUp("/a/b", 1))
	}
}
