package image

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/system/debian/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func DirectoryLink(
	v *semver.Version,
	architecture string,
) string {
	return locator.New(constant.Image).Path(
		"/cdimage/release/%s/%s/iso-cd",
		v.String(),
		architecture,
	).String()
}
