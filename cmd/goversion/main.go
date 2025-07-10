package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version"
	"github.com/funtimecoding/go-library/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.NotationArgument()
	monitor.AllArgument()
	pflag.String(
		argument.Skip,
		"",
		"Directory skip matches, comma-separated",
	)
	argument.ParseBind()
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	version.Check(
		argument.PositionalFallback(0, "."),
		argument.StringSlice(argument.Skip),
		runtime.ExecutableVersion().String(),
	)
	// TODO: Use option struct
	// TODO: Use REPOSITORY_PATH like gogitstatus
	// TODO: notation output and integrate into monitor
}
