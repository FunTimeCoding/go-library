package debian

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/debian/constant"
	"github.com/funtimecoding/go-library/pkg/system/debian/image"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func CheckLatestImage() {
	c := web.Client()

	for _, i := range strings.DeleteDuplicates(
		image.FindNames(
			web.GetString(
				c,
				locator.New(constant.Image).Path(
					"/cdimage/release/current/arm64/iso-cd",
				).Trail().String(),
			),
		),
	) {
		fmt.Printf("Image: %s\n", i)
	}
}
