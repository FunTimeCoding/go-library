//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert/fixture"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	system "github.com/funtimecoding/go-library/pkg/system/constant"
	"os"
	"path/filepath"
	"testing"
)

func openTestStore(t *testing.T) (*Store, *ollama.Client) {
	t.Helper()
	path := filepath.Join(t.TempDir(), constant.TestDatabase)

	return New(path), ollama.NewEnvironment()
}

func writeFixture(
	t *testing.T,
	directory string,
	name string,
	content string,
) {
	t.Helper()
	path := filepath.Join(directory, name)
	errors.PanicOnError(os.MkdirAll(filepath.Dir(path), 0o755))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o644))
}

func indexedTestStore(t *testing.T) (*Store, *ollama.Client) {
	t.Helper()
	s, o := openTestStore(t)
	s.AddCollection("test", fixture.Path(system.SearchPath), "**/*.md")
	s.Index("test")

	return s, o
}
