package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/macos/check/brew"
	"github.com/funtimecoding/go-library/pkg/macos/check/brew/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func main() {
	monitor.AllArgument()
	monitor.NotationArgument()
	monitor.VerboseArgument()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Verbose = viper.GetBool(argument.Verbose)
	brew.Check(o)
}
