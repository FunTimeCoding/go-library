package goversion

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/git/check/status"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goversion/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	programVersion string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, programVersion)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.String(argument.Skip, "", "Skip matches")
	pflag.Int(
		argument.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default: 3",
			item.GoVersion.Plural,
		),
	)
	monitor.ParseBind(programVersion, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Path = argument.PositionalFallback(
		0,
		environment.Fallback(
			status.RepositoryRootEnvironment,
			library.CurrentDirectory,
		),
	)
	o.Depth = viper.GetInt(argument.Depth)

	if s := environment.Optional(version.SkipEnvironment); s != "" {
		o.Skip = split.Comma(s)
	}

	if len(o.Skip) == 0 {
		o.Skip = argument.Slice(argument.Skip)
	}

	v := runtime.ExecutableVersion()

	if v == nil {
		system.Exitf(1, "could not get Go version\n")

		return
	}

	o.RuntimeVersion = v.String()
	version.Check(o)
}
