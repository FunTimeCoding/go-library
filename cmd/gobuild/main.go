package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.StringP(
		argument.Output,
		"o",
		"",
		"Output path",
	)
	pflag.StringP(
		argument.BuildTags,
		"t",
		"",
		"Build tags",
	)
	argument.ParseAndBind()
	build.Go(
		argument.Positional(0),
		viper.GetString(argument.Output),
		viper.GetString(argument.BuildTags),
	)
}
