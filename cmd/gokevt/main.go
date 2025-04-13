package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.Bool(
		argument.Clean,
		false,
		"Delete events older than 7 days",
	)
	argument.ParseAndBind()
	o := option.New()
	o.All = viper.GetBool(argument.All)
	o.Notation = viper.GetBool(argument.Notation)
	o.Clean = viper.GetBool(argument.Clean)
	event.Print(o)
}
