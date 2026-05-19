package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	git "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	assert.Strings(
		t,
		[]string{"hypertext", "markdown", "memory", "notation", "search", "wiki"},
		DirectoryContent(
			join.Absolute(
				FindDirectoryUp(WorkingDirectory(), git.Directory),
				constant.FixturePath,
			),
		),
	)
}
