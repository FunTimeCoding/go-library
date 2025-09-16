package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/git/check/status"
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	pflag.String(
		argument.Path,
		"",
		"Path to scan for git repositories. If not set, the current working directory will be used.",
	)
	pflag.Int(
		argument.Depth,
		3,
		fmt.Sprintf(
			"Depth to scan for %s. Default is 3.",
			constant.GoGitStatus.Plural,
		),
	)
	argument.ParseBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Verbose = viper.GetBool(argument.Verbose)
	o.Path = viper.GetString(argument.Path)
	o.Depth = viper.GetInt(argument.Depth)

	if s := environment.Default(
		status.RepositoryRootEnvironment,
		"",
	); s != "" {
		o.Path = s
	}

	status.Check(o)
}
