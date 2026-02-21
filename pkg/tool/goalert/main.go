package goalert

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/check/alert/option"
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
	pflag.Bool(argument.Critical, false, "Critical severity only")
	pflag.Bool(argument.Warning, false, "Warning severity only")
	pflag.Bool(argument.Extended, false, "Extended output")
	pflag.Bool(argument.Suppressed, false, "Include suppressed")
	pflag.Bool(argument.Old, false, "Include >1 week old")
	pflag.Bool(argument.Rules, false, "Print rules")
	pflag.Bool(argument.Firing, false, "Print firing rules")
	pflag.Bool(argument.Fingerprint, false, "Fingerprint column")
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Notation = viper.GetBool(argument.Notation)
	o.All = viper.GetBool(argument.All)
	o.Critical = viper.GetBool(argument.Critical)
	o.Warning = viper.GetBool(argument.Warning)
	o.Extended = viper.GetBool(argument.Extended)
	o.Old = viper.GetBool(argument.Old)
	o.Suppressed = viper.GetBool(argument.Suppressed)
	o.Rules = viper.GetBool(argument.Rules)
	o.Firing = viper.GetBool(argument.Firing)
	o.Fingerprint = viper.GetBool(argument.Fingerprint)
	o.Copyable = viper.GetBool(argument.Copyable)
	alert.Check(o)
}
