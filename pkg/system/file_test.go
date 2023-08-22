package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFile(t *testing.T) {
	SaveFile(Join(WorkingDirectory(), "test.txt"), "test content")
	assert.String(
		t,
		"test content",
		ReadFile(Join(WorkingDirectory(), "test.txt")),
	)
	DeleteFile("test.txt")
}
