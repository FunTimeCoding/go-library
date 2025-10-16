package git

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/github/action"
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestBranch(t *testing.T) {
	actual := Branch(system.ParentDirectory(constant.Depth))
	var e string

	if actual != constant.MainBranch {
		system.PrintEnvironment()
	}

	if action.IsActionRun() {
		e = constant.HeadReference
	} else {
		e = constant.MainBranch
	}

	assert.String(t, e, actual)
}
