package gokevt

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.Bool(
		argument.Clean,
		false,
		"Delete events older than 7 days",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Clean = viper.GetBool(argument.Clean)
	event.Print(o)
}
