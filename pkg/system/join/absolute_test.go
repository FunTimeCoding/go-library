package join

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAbsolute(t *testing.T) {
	assert.String(t, "/a/b/c", Absolute("a", "b", "c"))
	assert.String(t, "/a/b/c", Absolute("/a", "b", "c"))
	assert.String(t, "/a/b/c", Absolute("a", "/b", "c"))
	assert.String(t, "/a/b/c", Absolute("a", "b", "/c"))
	assert.String(t, "/a/b/c", Absolute("/a", "/b", "c"))
	assert.String(t, "/a/b/c", Absolute("a", "/b", "/c"))
	assert.String(t, "/a/b/c", Absolute("/a", "b", "/c"))
	assert.String(t, "/a/b/c", Absolute("/a/", "/b/", "/c/"))
}
