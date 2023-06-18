package git

import (
	"github.com/funtimecoding/go-library/system"
	"testing"
)

func TestStatus(t *testing.T) {
	Status(system.ParentDirectory(1))
}
