package run

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"runtime"
	"testing"
)

func TestRun(t *testing.T) {
	r1 := New()
	assert.True(t, r1.Panic)
	r1.Panic = false
	r1.Start("echo", "test")
	assert.True(t, r1.Error == nil)
	assert.String(t, "test\n", r1.OutputString)
	assert.String(t, "", r1.ErrorString)

	r2 := New()
	r2.Panic = false
	r2.Start("nonexistent")
	assert.True(t, r2.Error != nil)
	assert.String(t, "", r2.OutputString)
	assert.String(t, "", r2.ErrorString)

	switch runtime.GOOS {
	case constant.Windows:
		assert.String(
			t,
			`exec: "nonexistent": executable file not found in %PATH%`,
			r2.Error.Error(),
		)
	default:
		assert.String(
			t,
			`exec: "nonexistent": executable file not found in $PATH`,
			r2.Error.Error(),
		)
	}
}
