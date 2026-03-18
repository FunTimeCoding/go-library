package gosentry

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue/option"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gosentry", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	p := option.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Verbose = viper.GetBool(argument.Verbose)
	p.Copyable = viper.GetBool(argument.Copyable)
	issue.Check(p)
}
