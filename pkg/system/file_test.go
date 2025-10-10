package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"testing"
)

func TestFile(t *testing.T) {
	path := join.Absolute(WorkingDirectory(), "test.txt")
	SaveFile(path, "test content")
	assert.String(
		t,
		"test content",
		ReadFile(path),
	)
	DeleteFile("test.txt")
}
