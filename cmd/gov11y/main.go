package main

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/vulnerability/check/vulnerability"
	"github.com/funtimecoding/go-library/pkg/vulnerability/check/vulnerability/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String(argument.Filter, "", "modules, comma separated")
	monitor.VerboseArgument()
	argument.ParseBind()
	o := option.New()
	o.Verbose = viper.GetBool(argument.Verbose)
	o.Filter = argument.Slice(argument.Filter)
	vulnerability.Check(o)
}
