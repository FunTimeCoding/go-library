package git

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/github/action"
	github "github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/strings/contains"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"testing"
)

func TestBranch(t *testing.T) {
	system.PrintEnvironment()
	e := constant.MainBranch

	if action.IsActionRun() {
		if r := environment.Required(
			github.ReferenceEnvironment,
		); r != constant.MainBranch {
			e = r
		}
	}

	actual := Branch(system.ParentDirectory(constant.Depth))

	if false {
		// Sometimes HEAD
		assert.String(t, e, actual)
	}

	// TODO: Add reference environment to list if missing
	//  Then count somewhere what observations there are
	assert.True(
		t,
		contains.Any(
			[]string{actual},
			[]string{
				constant.MainBranch,
				constant.HeadReference,
			},
		),
	)
}
