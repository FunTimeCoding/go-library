package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	git "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	fixture := Join(
		FindDirectoryUp(WorkingDirectory(), git.Directory),
		constant.FixturePath,
	)
	assert.Strings(
		t,
		[]string{"hypertext", "lint", "markdown", "notation", "wiki"},
		DirectoryContent(fixture),
	)
}
