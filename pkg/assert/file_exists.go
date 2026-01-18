package assert

import (
	"os"
	"testing"
)

func FileExists(
	t *testing.T,
	path string,
) {
	if _, e := os.Stat(path); os.IsNotExist(e) {
		t.Fatalf("missing: %s", path)
	}
}
