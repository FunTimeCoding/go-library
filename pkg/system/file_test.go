package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"testing"
)

func TestFile(t *testing.T) {
	SaveFile(
		join.Absolute(WorkingDirectory(), "test.txt"),
		"test content",
	)
	assert.String(
		t,
		"test content",
		ReadFile(WorkingDirectory(), "test.txt"),
	)
	DeleteFile("test.txt")
}
