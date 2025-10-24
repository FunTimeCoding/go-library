package golint

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/lint"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.Bool(
		argument.Fix,
		false,
		"Fix concerns that can be fixed",
	)
	pflag.String(
		argument.Skip,
		"",
		"Directories to skip, comma separated",
	)
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	lint.Lint(
		viper.GetString(argument.Skip),
		viper.GetBool(argument.Verbose),
		viper.GetBool(argument.Fix),
	)
}
