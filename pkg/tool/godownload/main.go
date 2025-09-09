package godownload

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	common.Arguments()
	pflag.String(
		argument.PackageVersion,
		constant.LatestVersion,
		"Version to download, falls back to latest if not found",
	)
	pflag.String(
		argument.Output,
		environment.GetDefault("OUTPUT", download.DefaultOutput),
		"Output directory for executable",
	)
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Host = viper.GetString(argument.Host)
	o.Token = viper.GetString(argument.Token)
	o.Owner = viper.GetString(argument.Owner)
	o.Repository = viper.GetString(argument.Repository)
	o.PackageVersion = viper.GetString(argument.PackageVersion)
	o.Output = viper.GetString(argument.Output)
	o.Verbose = viper.GetBool(argument.Verbose)
	o.Package = argument.RequiredPositional(0, "PACKAGE")
	download.Run(o)
}
