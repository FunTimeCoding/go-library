package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"runtime"
	"testing"
)

func TestAbsoluteName(t *testing.T) {
	switch p := runtime.GOOS; p {
	case Linux:
		assert.True(t, AbsoluteName() != "")
	}
}
