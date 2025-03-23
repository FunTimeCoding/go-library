package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	RulesArgument  = "rules"
	FiringArgument = "firing"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.Bool(argument.Critical, false, "Critical severity only")
	pflag.Bool(argument.Warning, false, "Warning severity only")
	pflag.Bool(argument.Extended, false, "Extended output")
	pflag.Bool(argument.Suppressed, false, "Include suppressed")
	pflag.Bool(
		argument.Old,
		false,
		"Include alerts older than 1 week",
	)
	pflag.Bool(RulesArgument, false, "Print rules")
	pflag.Bool(FiringArgument, false, "Print firing rules")
	argument.ParseAndBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Critical = viper.GetBool(argument.Critical)
	o.Warning = viper.GetBool(argument.Warning)
	o.Extended = viper.GetBool(argument.Extended)
	o.Old = viper.GetBool(argument.Old)
	o.Suppressed = viper.GetBool(argument.Suppressed)
	o.Rules = viper.GetBool(RulesArgument)
	o.Firing = viper.GetBool(FiringArgument)
	alert.Print(o)
}
