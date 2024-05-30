package git

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.String(
		t,
		"main",
		Branch(system.ParentDirectory(constant.Depth)),
	)
}
