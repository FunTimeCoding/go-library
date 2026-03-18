package gochk

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gochk/check"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gochk", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	check.Check()
}
