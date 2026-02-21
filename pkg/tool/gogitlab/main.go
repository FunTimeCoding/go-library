package gogitlab

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Verbose = viper.GetBool(argument.Verbose)
	o.Copyable = viper.GetBool(argument.Copyable)
	job.Check(o)
}
