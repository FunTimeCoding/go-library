package file

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/fixture"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestFile(t *testing.T) {
	assert.NotNil(t, New(nil))
}

func TestParse(t *testing.T) {
	s := "## Example\n\nList of files in the current directory:\n```sh\nls -alh\n```\n"
	source := []byte(s)
	f := New(&source)
	assert.Strings(
		t,
		[]string{
			"Example",
			"List of files in the current directory:",
			"ls -alh\n",
		},
		f.Parse().Content(),
	)
}

func TestFixture(t *testing.T) {
	assert.String(
		t,
		"## Example\n\nList of files in the current directory:\n```sh\nls -alh\n```\n",
		string(system.ReadAll(fixture.File(constant.MarkdownPath, "1.md"))),
	)
}
