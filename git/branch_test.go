package git

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/system"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.String(
		t,
		"main",
		Branch(system.ParentDirectory(1)),
	)
}
