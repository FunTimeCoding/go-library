package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/outdated"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/outdated/option"
	"github.com/spf13/viper"
)

func main() {
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	argument.ParseBind()
	o := option.New()
	o.Copyable = viper.GetBool(argument.Copyable)
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	outdated.Check(o)
}
