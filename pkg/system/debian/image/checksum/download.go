package checksum

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
)

func Download(
	v *semver.Version,
	architecture string,
	workDirectory string,
) {
	path := Path(workDirectory)

	if system.FileExists(path) {
		fmt.Printf("Check sum exists: %s\n", path)

		return
	}

	l := Link(v, architecture)
	fmt.Printf("Check sum link: %s\n", l)
	web.Download(l, path)
}
