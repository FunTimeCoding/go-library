package build

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
)

func OutputFilename(v *semver.Version) string {
	return fmt.Sprintf("debian-%d.%d.img", v.Major, v.Minor)
}
