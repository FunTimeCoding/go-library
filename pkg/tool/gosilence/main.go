package gosilence

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/spf13/pflag"
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
	pflag.String(argument.Set, "", "Name, creates or updates")
	pflag.String(argument.Duration, "", "Duration, default 10m")
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Set = viper.GetString(argument.Set)
	o.Copyable = viper.GetBool(argument.Copyable)
	silence.Check(o)
}
