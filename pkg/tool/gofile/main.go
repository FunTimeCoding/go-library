package gofile

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/check/file"
	"github.com/funtimecoding/go-library/pkg/monitor/check/file/option"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Verbose = viper.GetBool(argument.Verbose)
	file.Check(o)
}
