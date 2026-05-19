//go:build local

package chunk

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestChunkGoSourceSplitsAtDeclarations(t *testing.T) {
	source := "package main\n\n" +
		"func Alpha() {\n" + strings.Repeat("\t// line\n", 200) + "}\n\n" +
		"func Beta() {\n" + strings.Repeat("\t// line\n", 200) + "}\n\n" +
		"func Gamma() {\n" + strings.Repeat("\t// line\n", 200) + "}\n"
	chunks := Document(source, "main.go")
	assert.Greater(t, 1, float64(len(chunks)))
	assert.StringContains(t, "package main", chunks[0].Text)
	found := false

	for _, c := range chunks {
		if strings.Contains(c.Text, "func Beta") {
			found = true
		}
	}

	if !found {
		t.Errorf("expected a chunk containing func Beta")
	}
}

func TestChunkGoSourceFallsBackOnParseError(t *testing.T) {
	invalid := "package main\n\nfunc broken {{{{\n" +
		strings.Repeat("text\n", 1000)
	chunks := Document(invalid, "broken.go")
	assert.Greater(t, 0, float64(len(chunks)))
}

func TestChunkGoSourceShortFile(t *testing.T) {
	source := "package store\n\nfunc Hello() {}\n"
	chunks := Document(source, "hello.go")
	assert.Count(t, 1, chunks)
	assert.String(t, source, chunks[0].Text)
}
