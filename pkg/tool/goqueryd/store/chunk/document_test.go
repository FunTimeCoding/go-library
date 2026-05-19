//go:build local

package chunk

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"strings"
	"testing"
)

func TestDocumentShortContent(t *testing.T) {
	content := "# Short\n\nThis is a short document.\n"
	chunks := Document(content, "test.md")
	assert.Count(t, 1, chunks)
	assert.String(t, content, chunks[0].Text)
	assert.Integer(t, 0, chunks[0].Position)
}

func TestDocumentSplitsAtHeadings(t *testing.T) {
	sections := make([]string, 5)

	for i := range sections {
		sections[i] = strings.Repeat("x", constant.ChunkSize/4)
	}

	content := "# First\n\n" + sections[0] +
		"\n# Second\n\n" + sections[1] +
		"\n# Third\n\n" + sections[2] +
		"\n# Fourth\n\n" + sections[3] +
		"\n# Fifth\n\n" + sections[4]
	chunks := Document(content, "test.md")
	assert.Greater(t, 1, float64(len(chunks)))
	assert.StringContains(t, "# First", chunks[0].Text)
}

func TestDocumentPreservesCodeFences(t *testing.T) {
	before := strings.Repeat("text ", constant.ChunkSize/8)
	code := strings.Repeat("line\n", 20)
	after := strings.Repeat("more ", constant.ChunkSize/8)
	content := "# Before\n\n" + before +
		"\n```go\n" + code + "```\n\n" +
		"# After\n\n" + after
	chunks := Document(content, "test.md")

	for _, c := range chunks {
		opens := strings.Count(c.Text, "```")

		if opens == 1 {
			t.Errorf("chunk splits inside code fence")
		}
	}
}
