package godebian

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/tool/godebian/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godebian/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.String(
		argument.Executable,
		"",
		"Executable to package: go-example",
	)
	pflag.String(
		argument.Version,
		"",
		"Package semantic version: 1.0.0, v-prefix gets trimmed",
	)
	pflag.String(
		maintainerNameArgument,
		"",
		"Maintainer name: AN Other",
	)
	pflag.String(
		maintainerEmailArgument,
		"",
		"Maintainer email: another@example.org",
	)
	pflag.Bool(
		systemdUnitFlag,
		false,
		"Create a systemd unit",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Executable = viper.GetString(argument.Executable)
	o.PackageVersion = semver.Trim(viper.GetString(argument.Version))
	o.MaintainerName = viper.GetString(maintainerNameArgument)
	o.MaintainerEmail = viper.GetString(maintainerEmailArgument)
	o.SystemdUnit = viper.GetBool(systemdUnitFlag)
	Run(o)
}
