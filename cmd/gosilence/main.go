package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/parameter"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.StringP(
		argument.Set,
		"s",
		"",
		"Create or update silence with name",
	)
	argument.ParseAndBind()
	p := parameter.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Set = viper.GetString(argument.Set)
	silence.Print(p)
}
