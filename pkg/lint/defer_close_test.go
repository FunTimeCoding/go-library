package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestDeferCloseDetected(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nfunc main() {\n\tdefer f.Close()\n}\n"),
	)
	assert.Boolean(t, true, r.HasConcerns())
}

func TestDeferCloseFixed(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nimport \"os\"\n\nfunc main() {\n\tdefer f.Close()\n}\n"),
	)
	assert.String(
		t,
		"package main\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/errors\"\n\t\"os\"\n)\n\nfunc main() {\n\tdefer errors.PanicClose(f)\n}\n",
		r.Fixed,
	)
}

func TestDeferCloseExistingImport(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nimport \"github.com/funtimecoding/go-library/pkg/errors\"\n\nfunc main() {\n\tdefer f.Close()\n}\n"),
	)
	assert.String(
		t,
		"package main\n\nimport \"github.com/funtimecoding/go-library/pkg/errors\"\n\nfunc main() {\n\tdefer errors.PanicClose(f)\n}\n",
		r.Fixed,
	)
}

func TestDeferCloseMultiImport(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nimport (\n\t\"os\"\n)\n\nfunc main() {\n\tdefer f.Close()\n}\n"),
	)
	assert.String(
		t,
		"package main\n\nimport (\n\t\"github.com/funtimecoding/go-library/pkg/errors\"\n\t\"os\"\n)\n\nfunc main() {\n\tdefer errors.PanicClose(f)\n}\n",
		r.Fixed,
	)
}

func TestDeferCloseNotTriggered(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nfunc main() {\n\tdefer errors.PanicClose(f)\n}\n"),
	)
	assert.Boolean(t, false, r.HasConcerns())
}

func TestDeferCloseNestedIndent(t *testing.T) {
	r := DeferClose(
		"test.go",
		strings.NewReader("package main\n\nimport \"os\"\n\nfunc main() {\n\t\tdefer f.Close()\n}\n"),
	)
	assert.StringContains(t, "\t\tdefer errors.PanicClose(f)", r.Fixed)
}
