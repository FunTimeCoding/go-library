package git

import (
	"github.com/funtimecoding/go-library/system"
	"testing"
)

func TestTags(t *testing.T) {
	Tags(system.ParentDirectory(1))
}
