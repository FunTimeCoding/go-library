package gov11y

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/vulnerability/check/vulnerability"
	"github.com/funtimecoding/go-library/pkg/vulnerability/check/vulnerability/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.String(argument.Filter, "", "modules, comma separated")
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Verbose = viper.GetBool(argument.Verbose)
	o.Filter = argument.Slice(argument.Filter)
	vulnerability.Check(o)
}
