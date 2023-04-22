package git

import (
	"github.com/funtimecoding/go-library/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestBranch(t *testing.T) {
	d, fail := os.Getwd()
	assert.FatalOnError(t, fail)

	b := Branch(filepath.Join(d, ".."))
	assert.String(t, "main", b)
}
