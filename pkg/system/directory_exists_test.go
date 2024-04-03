package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDirectoryExists(t *testing.T) {
	assert.True(t, DirectoryExists(WorkingDirectory()))
	assert.False(t, DirectoryExists("non-existent-directory"))
}
