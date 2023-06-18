package git

import (
	"github.com/funtimecoding/go-library/system"
	"path/filepath"
	"testing"
)

func TestTags(t *testing.T) {
	Tags(filepath.Join(system.WorkingDirectory(), ".."))
}
