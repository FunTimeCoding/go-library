package system

import (
	"archive/tar"
	"github.com/funtimecoding/go-library/pkg/integers64"
	"os"
)

func HeaderFileMode(h *tar.Header) os.FileMode {
	return os.FileMode(integers64.ToUnsigned32(h.Mode))
}
