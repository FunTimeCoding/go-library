package gosalt

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("gosalt", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	// TODO: Use Salt API, allow calling it from CI/CD
}
