package runtime

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/constant"
	"runtime"
	"strings"
)

func Version() *semver.Version {
	return semver.New(strings.TrimPrefix(runtime.Version(), constant.Go))
}
