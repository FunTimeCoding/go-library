package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	git "github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"testing"
)

func TestDirectoryContent(t *testing.T) {
	fixture := join.Absolute(
		FindDirectoryUp(WorkingDirectory(), git.Directory),
		constant.FixturePath,
	)
	fmt.Printf("Fixture: %s\n", fixture)
	assert.Strings(
		t,
		[]string{"hypertext", "lint", "markdown", "notation", "wiki"},
		DirectoryContent(fixture),
	)
}
