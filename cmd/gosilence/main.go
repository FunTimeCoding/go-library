package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.String(
		argument.Set,
		"",
		"Create or update silence with name",
	)
	argument.ParseAndBind()
	p := option.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Set = viper.GetString(argument.Set)
	silence.Print(p)
}
