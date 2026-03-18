package naming

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSegmentsCamelCase(t *testing.T) {
	assert.Strings(t, []string{"dir", "name"}, segments("dirName"))
}

func TestSegmentsPascalCase(t *testing.T) {
	assert.Strings(t, []string{"dir", "something"}, segments("DirSomething"))
}

func TestSegmentsSnakeCase(t *testing.T) {
	assert.Strings(t, []string{"model", "context"}, segments("model_context"))
}

func TestSegmentsSingleWord(t *testing.T) {
	assert.Strings(t, []string{"directory"}, segments("directory"))
}

func TestSegmentsAllLower(t *testing.T) {
	assert.Strings(t, []string{"url"}, segments("url"))
}

func TestSegmentsMixed(t *testing.T) {
	assert.Strings(t, []string{"output", "directory"}, segments("OutputDirectory"))
}
