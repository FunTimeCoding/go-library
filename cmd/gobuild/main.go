package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String(argument.Output, "", "Output path")
	argument.ParseAndBind()
	build.Go(
		argument.Positional(0),
		viper.GetString(argument.Output),
	)
}
