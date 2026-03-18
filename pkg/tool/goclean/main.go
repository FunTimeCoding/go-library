package goclean

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
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
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goclean", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.GitLabHost = environment.Required(constant.HostEnvironment)
	o.Verbose = viper.GetBool(argument.Verbose)
	clean.Run(o)
}
