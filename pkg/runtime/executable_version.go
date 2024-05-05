package runtime

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"regexp"
)

func ExecutableVersion() *semver.Version {
	m := regexp.MustCompile(
		`go(\d+\.\d+\.\d+)`,
	).FindStringSubmatch(system.Run(constant.Go, constant.Version))

	if len(m) > 1 {
		return semver.New(m[1])
	}

	return nil
}
