package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitConstant "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	fixture := Join(
		FindDirectoryUp(
			WorkingDirectory(),
			gitConstant.Directory,
		),
		constant.FixturePath,
	)
	assert.Strings(t, []string{"lint", "notation"}, DirectoryContent(fixture))
}
