package goghexporter

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goghexporter/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.VerboseArgument()
	pflag.String(
		argument.Owner,
		environment.Optional(constant.OwnerEnvironment),
		"GitHub owner",
	)
	pflag.Duration(argument.Interval, 5*time.Minute, "Poll interval")
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Owner = viper.GetString(argument.Owner)
	o.Interval = viper.GetDuration(argument.Interval)
	o.Verbose = viper.GetBool(argument.Verbose)
	Run(o, r)
}
