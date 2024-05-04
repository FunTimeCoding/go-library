package build

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	build := New("a", "b", "c")
	assert.String(t, build.Version, "a")
	assert.String(t, build.GitHash, "b")
	assert.String(t, build.BuildDate, "c")
	assert.True(t, len(Date()) > 0)
	assert.True(t, len(GitHash()) > 0)
	assert.True(t, len(GitTag()) > 0)
}
