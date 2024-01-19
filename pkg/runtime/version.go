package runtime

import (
	"github.com/coreos/go-semver/semver"
	"runtime"
	"strings"
)

func Version() *semver.Version {
	return semver.New(strings.TrimPrefix(runtime.Version(), "go"))
}
