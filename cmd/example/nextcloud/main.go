package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/nextcloud"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func main() {
	n := nextcloud.NewEnvironment()

	if false {
		n.Status()
	}

	if false {
		for _, f := range n.ReadDirectory("/") {
			fmt.Printf(
				"Name: %s, IsDir: %t, Size: %d\n",
				f.Name(),
				f.IsDir(),
				f.Size(),
			)
		}
	}

	if false {
		n.DownloadFile(
			"example.png",
			join.Absolute(
				system.Home(),
				constant.DownloadsPath,
				"example.png",
			),
		)
	}

	if false {
		n.UploadFile(
			"example2.png",
			join.Absolute(
				system.Home(),
				constant.DownloadsPath,
				"example.png",
			),
		)
	}
}
