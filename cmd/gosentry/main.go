package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue"
	"github.com/funtimecoding/go-library/pkg/sentry/check/issue/option"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	argument.ParseBind()
	p := option.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Verbose = viper.GetBool(argument.Verbose)
	issue.Print(p)
}
