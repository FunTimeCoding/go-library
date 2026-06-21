package testutil

import (
	"os"
	"testing"
)

func ReadFile(
	t *testing.T,
	path string,
) string {
	t.Helper()
	b, e := os.ReadFile(path)

	if e != nil {
		t.Fatal(e)
	}

	return string(b)
}
