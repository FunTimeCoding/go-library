package git

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/system"
	"path/filepath"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.String(
		t,
		"main",
		Branch(filepath.Join(system.WorkingDirectory(), "..")),
	)
}
