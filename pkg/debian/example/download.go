package example

import (
	"github.com/funtimecoding/go-library/pkg/debian"
	"runtime"
)

func Download() {
	debian.CheckLatestImage()
	debian.New().DownloadImage(debian.Bookworm, runtime.GOARCH)
}
