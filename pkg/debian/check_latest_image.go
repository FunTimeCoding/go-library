package debian

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/debian/image"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web"
)

func CheckLatestImage() {
	c := web.Client(true)

	for _, element := range strings.DeleteDuplicates(
		image.FindNames(
			web.GetString(
				c,
				"https://cdimage.debian.org/cdimage/release/current/arm64/iso-cd/",
			),
		),
	) {
		fmt.Printf("Image: %s\n", element)
	}
}
