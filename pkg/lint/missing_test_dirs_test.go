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
	result := missingTestDirs(paths, nil)
	assert.Integer(t, 1, len(result))
	_, ok := result["pkg/foo"]
	assert.Boolean(t, true, ok)
}

func TestMissingTestDirsNone(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
		"pkg/foo/foo_test.go",
	}
	result := missingTestDirs(paths, nil)
	assert.Integer(t, 0, len(result))
}

func TestMissingTestDirsGenerated(t *testing.T) {
	paths := []string{
		"pkg/foo/foo.go",
	}
	generatedPaths := []string{
		"pkg/baz/generated.go",
	}
	result := missingTestDirs(paths, generatedPaths)
	assert.Integer(t, 2, len(result))
	_, ok := result["pkg/foo"]
	assert.Boolean(t, true, ok)
	_, ok = result["pkg/baz"]
	assert.Boolean(t, true, ok)
}
