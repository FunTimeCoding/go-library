package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/system/argument"
	viperHelper "github.com/funtimecoding/go-library/pkg/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String(argument.Main, ".", "Main path")
	pflag.String(argument.Output, "", "Output path")
	viperHelper.ParseAndBind()
	viperHelper.RequiredString(argument.Output)

	if s := build.Go(
		viper.GetString(argument.Main),
		viper.GetString(argument.Output),
	); s != "" {
		fmt.Printf("Output:\n%s", s)
	}
}
