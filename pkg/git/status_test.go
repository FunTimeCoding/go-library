package git

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"testing"
)

func TestStatus(t *testing.T) {
	Status(system.ParentDirectory(depth))
}
