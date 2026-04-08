package gomatmcp

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	gomonitor "github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gomatmcp", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	gomonitor.ParseBind(version, gitHash, buildDate)
	Run()
}
