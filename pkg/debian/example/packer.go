package example

import (
	"github.com/funtimecoding/go-library/pkg/debian"
	"runtime"
)

func Packer() {
	debian.New().Packer(debian.Bookworm, runtime.GOARCH)
}
