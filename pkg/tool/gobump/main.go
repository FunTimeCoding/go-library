package gobump

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gobump/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gobump", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Increase = argument.RequiredPositional(0, "INCREASE")
	Run(o)
}
