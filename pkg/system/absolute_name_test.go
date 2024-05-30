package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"runtime"
	"testing"
)

func TestAbsoluteName(t *testing.T) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		assert.True(t, AbsoluteName() != "")
	}
}
