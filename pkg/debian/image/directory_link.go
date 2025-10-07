package image

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
)

func DirectoryLink(
	v *semver.Version,
	architecture string,
) string {
	return fmt.Sprintf(
		"https://cdimage.debian.org/cdimage/release/%s/%s/iso-cd",
		v.String(),
		architecture,
	)
}
