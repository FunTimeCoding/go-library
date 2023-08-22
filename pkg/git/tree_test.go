package git

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestTree(t *testing.T) {
	Tree(system.ParentDirectory(1))
}
