package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWindowsToLinuxPath(t *testing.T) {
	assert.String(
		t,
		"/c/Users/example/Desktop/file.txt",
		WindowsToLinuxPath("C:\\Users\\example\\Desktop\\file.txt"),
	)
}
