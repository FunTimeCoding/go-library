package goclean

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	gitlab "github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/constant"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.GitLabHost = environment.Required(gitlab.HostEnvironment)
	o.Verbose = viper.GetBool(argument.Verbose)
	clean.Run(o)
}
