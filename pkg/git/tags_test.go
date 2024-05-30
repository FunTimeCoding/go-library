package git

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestTags(t *testing.T) {
	Tags(system.ParentDirectory(constant.Depth))
}
