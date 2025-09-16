package goclean

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.GitLabHost = environment.Exit(constant.HostEnvironment)
	o.Verbose = viper.GetBool(argument.Verbose)
	clean.Run(o)
}
