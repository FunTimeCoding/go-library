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
	var e string

	if action.IsActionRun() {
		e = environment.Required(github.ReferenceEnvironment)
	} else {
		e = constant.MainBranch
	}

	assert.String(t, e, Branch(system.ParentDirectory(constant.Depth)))
}
