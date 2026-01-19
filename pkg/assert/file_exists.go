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
		t.Helper()
		t.Errorf("missing: %s", path)
	}
}
