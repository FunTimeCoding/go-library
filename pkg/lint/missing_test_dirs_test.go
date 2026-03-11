package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMissingTestDirs(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
		"pkg/bar/bar.go",
		"pkg/bar/bar_test.go",
	}
	result := missingTestDirs(paths)
	assert.Integer(t, 1, len(result))
	_, ok := result["pkg/foo"]
	assert.Boolean(t, true, ok)
}

func TestMissingTestDirsNone(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
		"pkg/foo/foo_test.go",
	}
	result := missingTestDirs(paths)
	assert.Integer(t, 0, len(result))
}
