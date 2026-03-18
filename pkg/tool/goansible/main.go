package goansible

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goansible", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	// TODO: Find Ansible API, allow calling it from CI/CD
}
