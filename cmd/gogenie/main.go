package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	argument.ParseAndBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	alert.Print(o)
}
