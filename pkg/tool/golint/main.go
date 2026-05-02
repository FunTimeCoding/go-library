package golint

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lint"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("golint", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.Bool(
		argument.Fix,
		false,
		"Fix concerns that can be fixed",
	)
	pflag.String(
		argument.Skip,
		"",
		"Directories to skip, comma separated",
	)
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	lint.Lint(
		viper.GetString(argument.Skip),
		viper.GetBool(argument.Verbose),
		viper.GetBool(argument.Fix),
	)
}
