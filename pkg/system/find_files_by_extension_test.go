package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitConstant "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestFindFilesByExtension(t *testing.T) {
	fixture := Join(
		FindDirectoryUp(
			WorkingDirectory(),
			gitConstant.Directory,
		),
		constant.FixturePath,
	)
	assert.Count(
		t,
		2,
		FindFilesByExtensions(fixture, ".json"),
	)
}
