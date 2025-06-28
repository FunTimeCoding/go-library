package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	argument.ParseBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Verbose = viper.GetBool(argument.Verbose)
	job.Check(o)
}
