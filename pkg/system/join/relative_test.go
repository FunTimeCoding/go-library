package join

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRelative(t *testing.T) {
	assert.String(t, "a/b/c", Relative("a", "b", "c"))
	assert.String(t, "a/b/c", Relative("/a", "b", "c"))
	assert.String(t, "a/b/c", Relative("a", "/b", "c"))
	assert.String(t, "a/b/c", Relative("a", "b", "/c"))
	assert.String(t, "a/b/c", Relative("/a", "/b", "c"))
	assert.String(t, "a/b/c", Relative("a", "/b", "/c"))
	assert.String(t, "a/b/c", Relative("/a", "b", "/c"))
	assert.String(t, "a/b/c", Relative("/a/", "/b/", "/c/"))
}
