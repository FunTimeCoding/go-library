package testutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func CopyTree(
	t *testing.T,
	source string,
	destination string,
) {
	t.Helper()
	absolute, e := filepath.Abs(source)

	if e != nil {
		t.Fatalf("abs: %s", e)
	}

	e = filepath.WalkDir(
		absolute,
		func(
			path string,
			d fs.DirEntry,
			walkError error,
		) error {
			if walkError != nil {
				return walkError
			}

			relative, e := filepath.Rel(absolute, path)

			if e != nil {
				return e
			}

			target := filepath.Join(destination, relative)

			if d.IsDir() {
				return os.MkdirAll(target, 0755)
			}

			content, e := os.ReadFile(path)

			if e != nil {
				return e
			}

			return os.WriteFile(target, content, 0644)
		},
	)

	if e != nil {
		t.Fatalf("copy tree: %s", e)
	}
}
