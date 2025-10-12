package git

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/github/action"
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestBranch(t *testing.T) {
	var e string

	if action.IsActionRun() {
		e = constant.HeadReference
	} else {
		e = constant.MainBranch
	}

	assert.String(t, e, Branch(system.ParentDirectory(constant.Depth)))
}
