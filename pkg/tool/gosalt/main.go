package gosalt

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gosalt", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	// TODO: Use Salt API, allow calling it from CI/CD
}
