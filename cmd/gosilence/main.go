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
	pflag.String(argument.Set, "", "Name, creates or updates")
	pflag.String(argument.Duration, "", "Duration, default 10m")
	argument.ParseAndBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Set = viper.GetString(argument.Set)
	silence.Print(o)
}
