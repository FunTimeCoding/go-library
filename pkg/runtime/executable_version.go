package runtime

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/expression"
	"github.com/funtimecoding/go-library/pkg/system"
)

func ExecutableVersion() *semver.Version {
	if m := expression.SubMatch(
		`go(\d+\.\d+\.\d+)`,
		system.Run(constant.Go, constant.Version),
	); m != "" {
		return semver.New(m)
	}

	return nil
}
