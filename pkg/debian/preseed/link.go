package preseed

import (
	"github.com/funtimecoding/go-library/pkg/debian/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func Link(release string) string {
	return locator.New(constant.Web).Path(
		"/releases/%s/example-preseed.txt",
		release,
	).String()
}
