package image

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"github.com/funtimecoding/go-library/pkg/web"
)

func Download(
	v *semver.Version,
	architecture string,
	workDirectory string,
) {
	i := Name(v, architecture)
	path := join.Absolute(workDirectory, i)

	if system.FileExists(path) {
		fmt.Printf("Image exists: %s\n", path)

		return
	}

	l := Link(v, architecture, i)
	fmt.Printf("Image link: %s\n", l)
	web.Download(l, path)
}
