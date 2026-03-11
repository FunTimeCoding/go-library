package vfs

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestWriteRead(t *testing.T) {
	v := New()
	v.Write("a/b.yaml", "content")
	assert.String(t, "content", v.Read("a/b.yaml"))
}

func TestHas(t *testing.T) {
	v := New()
	v.Write("a.yaml", "x")
	assert.Boolean(t, true, v.Has("a.yaml"))
	assert.Boolean(t, false, v.Has("b.yaml"))
}

func TestDelete(t *testing.T) {
	v := New()
	v.Write("a.yaml", "x")
	v.Delete("a.yaml")
	assert.Boolean(t, false, v.Has("a.yaml"))
}

func TestFilesSorted(t *testing.T) {
	v := New()
	v.Write("b.yaml", "")
	v.Write("a.yaml", "")
	v.Write("c.yaml", "")
	assert.Any(t, []string{"a.yaml", "b.yaml", "c.yaml"}, v.Files())
}

func TestFlushFrom(t *testing.T) {
	v := New()
	v.Write("sub/file.txt", "hello")
	dir := t.TempDir()
	v.Flush(dir)
	back := From(dir)
	assert.String(t, "hello", back.Read("sub/file.txt"))
}
