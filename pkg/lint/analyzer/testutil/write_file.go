package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

func WriteFile(
	t *testing.T,
	base string,
	path string,
	content string,
) {
	t.Helper()
	full := filepath.Join(base, path)
	e := os.MkdirAll(filepath.Dir(full), 0755)

	if e != nil {
		t.Fatal(e)
	}

	e = os.WriteFile(full, []byte(content), 0644)

	if e != nil {
		t.Fatal(e)
	}
}
