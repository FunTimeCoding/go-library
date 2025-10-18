package gorunif

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.String(constant.Base, "HEAD~1", "Base commit")
	pflag.String(constant.Head, "HEAD", "Head commit")
	monitor.VerboseArgument()
	argument.ParseBind()
	o := option.New()
	o.Verbose = viper.GetBool(argument.Verbose)
	run_if.Run(o)
}
