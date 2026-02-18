package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func main() {
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	argument.ParseBind()
	p := option.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Verbose = viper.GetBool(argument.Verbose)
	p.Copyable = viper.GetBool(argument.Copyable)
	issue.Check(p)
}
