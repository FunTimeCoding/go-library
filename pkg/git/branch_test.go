package git

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/github/action"
	github "github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"testing"
)

func TestBranch(t *testing.T) {
	system.PrintEnvironment()
	e := constant.HeadReference

	if action.IsActionRun() {
		if r := environment.Required(
			github.ReferenceEnvironment,
		); r != constant.MainBranch {
			e = r
		}
	} else {
		e = constant.MainBranch
	}

	assert.String(t, e, Branch(system.ParentDirectory(constant.Depth)))
}
