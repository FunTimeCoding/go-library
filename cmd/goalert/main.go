package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/parameter"
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
	p := parameter.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Critical = viper.GetBool(argument.Critical)
	p.Warning = viper.GetBool(argument.Warning)
	p.Extended = viper.GetBool(argument.Extended)
	p.Old = viper.GetBool(argument.Old)
	p.Suppressed = viper.GetBool(argument.Suppressed)
	p.Rules = viper.GetBool(RulesArgument)
	p.Firing = viper.GetBool(FiringArgument)
	alert.Print(p)
}
