package file

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFile(t *testing.T) {
	SaveFile("test.txt", "test content")
	assert.String(t, "test content", ReadFile("test.txt"))
	DeleteFile("test.txt")
}
