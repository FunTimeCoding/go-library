package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/go_mod/version_check"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	var skip []string
	pflag.String(
		argument.Skip,
		"",
		"Directory skip matches, comma-separated",
	)
	argument.ParseBind()

	if s := viper.GetString(argument.Skip); s != "" {
		skip = split.Comma(s)
	}

	version_check.Scan(
		argument.PositionalFallback(0, "."),
		skip,
		runtime.ExecutableVersion().String(),
	)
}
