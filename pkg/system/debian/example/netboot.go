package example

import (
	"github.com/funtimecoding/go-library/pkg/system/debian"
	"runtime"
)

func Netboot() {
	debian.New().DownloadNetboot(debian.Bookworm, runtime.GOARCH)
}
