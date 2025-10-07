package example

import (
	"github.com/funtimecoding/go-library/pkg/debian"
	"runtime"
)

func Netboot() {
	debian.New().DownloadNetboot(debian.Bookworm, runtime.GOARCH)
}
