package goraidd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goraidd", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.String(
		argument.LogCache,
		"",
		"Path to LogDataCache.json",
	)
	pflag.String(
		argument.EliteInsights,
		"",
		"Path to Elite Insights output directory",
	)
	pflag.String(
		argument.Output,
		"",
		"Path for generated report files",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.LogCachePath = viper.GetString(argument.LogCache)
	o.EliteInsightsPath = viper.GetString(argument.EliteInsights)
	o.OutputPath = viper.GetString(argument.Output)
	Run(o)
}
