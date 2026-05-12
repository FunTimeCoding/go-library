package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func WriteModFile(
	t *testing.T,
	directory string,
	moduleName string,
	extraLines ...string,
) {
	t.Helper()
	modContent := fmt.Sprintf("module %s\n\ngo 1.26.3\n", moduleName)

	for _, line := range extraLines {
		modContent = fmt.Sprintf("%s\n%s\n", modContent, line)
	}

	e := os.WriteFile(
		filepath.Join(directory, "go.mod"),
		[]byte(modContent),
		0644,
	)

	if e != nil {
		t.Fatalf("write go.mod: %s", e)
	}
}
