package virtual_file_system

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

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

func TestDirectoryExists(t *testing.T) {
	v := New()
	v.Write("pkg/tool/gotest/server/router.go", "")
	v.Write("pkg/tool/gotest/constant/constant.go", "")
	assert.Boolean(t, true, v.DirectoryExists("pkg/tool/gotest"))
	assert.Boolean(t, true, v.DirectoryExists("pkg/tool/gotest/server"))
	assert.Boolean(t, true, v.DirectoryExists("pkg"))
	assert.Boolean(
		t,
		false,
		v.DirectoryExists("pkg/tool/gotest/model_context"),
	)
	assert.Boolean(t, false, v.DirectoryExists("pkg/tool/other"))
}

func TestReadDirectory(t *testing.T) {
	v := New()
	v.Write("pkg/tool/gotest/server/router.go", "")
	v.Write("pkg/tool/gotest/server/handler.go", "")
	v.Write("pkg/tool/gotest/constant/constant.go", "")
	v.Write("pkg/tool/gotest/run.go", "")
	assert.Any(
		t,
		[]string{"constant", "run.go", "server"},
		v.MustReadDirectory("pkg/tool/gotest"),
	)
	assert.Any(
		t,
		[]string{"handler.go", "router.go"},
		v.MustReadDirectory("pkg/tool/gotest/server"),
	)
}

func TestReadDirectoryMissing(t *testing.T) {
	v := New()
	result, e := v.ReadDirectory("pkg/tool/missing")
	assert.True(t, result == nil)
	assert.NotNil(t, e)
	assert.True(t, errors.Is(e, ErrorDirectoryNotFound))
}

func TestFlushFrom(t *testing.T) {
	v := New()
	v.Write("sub/file.txt", "hello")
	d := t.TempDir()
	v.Flush(d)
	back := From(d)
	assert.String(t, "hello", back.Read("sub/file.txt"))
}
