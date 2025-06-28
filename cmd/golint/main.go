package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/lint"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Bool(
		argument.Fix,
		false,
		"Go concerns that can be fixed",
	)
	pflag.String(
		argument.Skip,
		"",
		"Directories to skip, comma separated",
	)
	monitor.VerboseArgument()
	argument.ParseBind()
	lint.Lint(
		viper.GetString(argument.Skip),
		viper.GetBool(argument.Verbose),
		viper.GetBool(argument.Fix),
	)
}
