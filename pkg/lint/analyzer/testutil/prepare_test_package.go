package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func PrepareTestPackage(
	t *testing.T,
	sourceDirectory string,
	extraModLines ...string,
) string {
	t.Helper()
	absolute, e := filepath.Abs(sourceDirectory)

	if e != nil {
		t.Fatalf("abs: %s", e)
	}

	temporary := t.TempDir()
	CopyTree(t, absolute, temporary)
	modContent := "module example\n\ngo 1.26.3\n"

	for _, line := range extraModLines {
		modContent = fmt.Sprintf("%s\n%s\n", modContent, line)
	}

	e = os.WriteFile(
		filepath.Join(temporary, "go.mod"),
		[]byte(modContent),
		0644,
	)

	if e != nil {
		t.Fatalf("write go.mod: %s", e)
	}

	return temporary
}
