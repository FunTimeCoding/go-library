package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	git "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestFindFilesByExtension(t *testing.T) {
	fixture := Join(
		FindDirectoryUp(WorkingDirectory(), git.Directory),
		constant.FixturePath,
	)
	assert.Count(
		t,
		5,
		FindFilesByExtension(fixture, ".json"),
	)
}
